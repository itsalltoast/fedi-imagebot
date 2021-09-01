package store

/*
 * fedi-imagebot: An imagebot for the Fediverse.
 * Copyright © 2021, Mick 🔥 Abernathy <@itsalltoast@to.ast.my>
 *   BSD-3 - See LICENSE for usage restrictions.
 */

import "github.com/itsalltoast/fedi-imagebot/config"

// Interface provides a standardized set of functions that
// can be performed on a data storage backend.
//
type Interface interface {
	HaveURL(url string) bool
	AddURL(url string) error
	AddURLs(url []string) error
	MarkURL(url string) error
	GetRandomURL(unseenOnly bool) (string, error)
	Close()
}

// Store provides a base data storage object that maintains a
// link back to the store's configuration object.
//
type Store struct {
	Interface

	cfg *config.Config
}

// NewStore creates the appropriate Store object for the user's
// configured data storage backend.
//
func NewStore(c *config.Config) Interface {
	if c.DBType == "sqlite3" {
		st := newSQLStore("sqlite3", c.DBConnectString)
		st.cfg = c
		return st
	}

	return nil
}
