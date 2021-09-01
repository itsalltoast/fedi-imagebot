/*
 * fedi-imagebot: An imagebot for the Fediverse.
 * Copyright © 2021, Mick 🔥 Abernathy <@itsalltoast@to.ast.my>
 *   BSD-3 - See LICENSE for usage restrictions.
 */
package store

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"strings"
)

type SQLStore struct {
	Store

	db   *sql.DB
	add  *sql.Stmt
	rand *sql.Stmt
	mark *sql.Stmt
}

func NewSQLStore(dbType string, connectString string) *SQLStore {
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

func (s *SQLStore) HaveURL(url string) bool {

	// TODO: NOT IMPLEMENTED but kind of doesn't matter since there's a unique index on bot_images.url anyway.
	//
	return false
}

func (s *SQLStore) AddURL(url string) error {
	_, e := s.addStmt().Exec(url)
	return e
}

func (s *SQLStore) AddURLs(urls []string) error {

	var e error
	for _, url := range urls {
		if e = s.AddURL(url); e != nil {
			return e
		}
	}

	return nil
}

func (s *SQLStore) MarkURL(url string) error {
	_, e := s.getMarkStmt().Exec(url)
	return e
}

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
