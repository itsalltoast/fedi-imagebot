/*
 * fedi-imagebot: An imagebot for the Fediverse.
 * Copyright © 2021, Mick 🔥 Abernathy <@itsalltoast@to.ast.my>
 *   BSD-3 - See LICENSE for usage restrictions.
 */
package store

import "github.com/itsalltoast/fedi-imagebot/config"

type StoreInterface interface {
	HaveURL(url string) bool
	AddURL(url string) error
	AddURLs(url []string) error
	MarkURL(url string) error
	GetRandomURL(unseenOnly bool) (string, error)
	Close()
}

type Store struct {
	StoreInterface

	cfg *config.Config
}

func NewStore(c *config.Config) StoreInterface {
	if c.DBType == "sqlite3" {
		st := NewSQLStore("sqlite3", c.DBConnectString)
		st.cfg = c
		return st
	}

	return nil
}
