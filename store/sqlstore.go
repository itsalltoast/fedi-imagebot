package store
/*
 * fedi-imagebot: An imagebot for the Fediverse.
 * Copyright © 2021, Mick 🔥 Abernathy <@itsalltoast@to.ast.my>
 *   BSD-3 - See LICENSE for usage restrictions.
 */

import (
	"database/sql"
	// Load the SQLite3 driver as a dependency.
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"strings"
)

// SQLStore is a backend that supports SQL-based databases (but currently
// is only really useful for SQLite3 files).
//
type SQLStore struct {
	Store

	db   *sql.DB
	add  *sql.Stmt
	rand *sql.Stmt
	mark *sql.Stmt
}

// newSQLStore creates a storage object pointing to the requested database
// configuration and returns it to the user.
//
func newSQLStore(dbType string, connectString string) *SQLStore {
	s := new(SQLStore)
	var e error
	if len(connectString) < 1 {
		connectString = "${HOME}/.fedi-imagebot.db"
	}
	connectString = os.ExpandEnv(connectString)
	s.db, e = sql.Open(dbType, connectString)
	if e != nil {
		log.Println("Failed opening database:", connectString)
	}
	return s
}

// Close an established SQLStore connection.
//
func (s *SQLStore) Close() {
	e := s.db.Close()
	if e != nil {
		log.Println("Failed closing database:", e)
	}
}

// TODO: SQL and statments are all as generic as possible but really only tested on sqlite3.  Support for other databases (for monster sites) incoming.
// TODO: Support multiple bots per database (currently limited to 1-1).
//
func (s *SQLStore) createTable() error {
	_, e := s.db.Exec(`create table bot_images (url string unique, discovered timestamp default current_timestamp, posted timestamp)`)
	return e
}

func (s *SQLStore) prepare(query string) *sql.Stmt {
	var e error
	var ret *sql.Stmt
	ret, e = s.db.Prepare(query)

	// The only time this should error out is if we mangled our SQL.
	//
	if e != nil && strings.HasPrefix(e.Error(), "no such table") {

		e = s.createTable()
		// We should not continue if this fails.
		if e != nil {
			log.Fatalln(s.cfg.Name, "Failed initializing database:", e)
		}

		// Even though this is recursive, it should only run once.  Either we succeed here or we fail there.
		//
		return s.prepare(query)

	} else if e != nil {
		// If we return nil or an uninitialied sql.Stmt, the program will crash hard?
		log.Fatalln(s.cfg.Name, "Failed preparing query:", e)
	}

	return ret
}

func (s *SQLStore) addStmt() *sql.Stmt {
	if s.add != nil {
		return s.add
	}

	s.add = s.prepare(`insert into bot_images (url) values (?)`)
	return s.add
}

func (s *SQLStore) getRandStmt(unseenOnly bool) *sql.Stmt {
	if s.rand != nil {
		return s.rand
	}

	if unseenOnly {
		// Only return a URL that has never been posted before.  Note that this does not guarantee that the
		// image is unique: If it's a copy of some other search result image, it'll look like a repost.  It
		// might be worth trying to store checksums down the road to prevent reposts in the future, but that's
		// still not perfect and is probably more trouble than it's worth.
		//
		s.rand = s.prepare(`select url from bot_images where posted is null order by random() limit 1`)
	} else {

		// TODO: We want to hint at some kind of time-preference to ensure we
		// are posting older images before repeating recent ones.
		//
		s.rand = s.prepare(`select url from bot_images order by random() limit 1`)
	}

	return s.rand
}

func (s *SQLStore) getMarkStmt() *sql.Stmt {
	if s.mark != nil {
		return s.mark
	}

	s.mark = s.prepare(`update bot_images set posted = current_timestamp where url = ?`)
	return s.mark
}

// HaveURL reports whether a URL already exists in the database or not.  This function may go away,
// since the database is expected to have a constraint that prohibits the addition of duplicate
// URLs.
//
func (s *SQLStore) HaveURL(url string) bool {

	// TODO: NOT IMPLEMENTED but kind of doesn't matter since there's a unique index on bot_images.url anyway.
	//
	return false
}

// AddURL inserts a new URL into the database.
//
func (s *SQLStore) AddURL(url string) error {
	_, e := s.addStmt().Exec(url)
	return e
}

// AddURLs adds a set of new URLs into the database (using store.AddURL).
//
func (s *SQLStore) AddURLs(urls []string) error {

	var e error
	for _, url := range urls {
		if e = s.AddURL(url); e != nil {
			return e
		}
	}

	return nil
}

// MarkURL updates the database to indicate that a URL has been posted to
// the bot's timeline.
//
func (s *SQLStore) MarkURL(url string) error {
	_, e := s.getMarkStmt().Exec(url)
	return e
}

// GetRandomURL returns a single random URL from the database.  If unseenOnly
// is true, it will exclude any URLs that have already been posted
// to the user's timeline.
//
func (s *SQLStore) GetRandomURL(unseenOnly bool) (string, error) {

	r, e := s.getRandStmt(unseenOnly).Query()
	if e != nil {
		return "", e
	}
	defer r.Close()

	if !r.Next() {
		return "", sql.ErrNoRows
	}

	var ret sql.NullString
	e = r.Scan(&ret)

	// If there was an error parsing the string, it should be returned here with a blank string.
	//
	return ret.String, e
}
