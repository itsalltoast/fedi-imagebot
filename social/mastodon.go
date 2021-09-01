package social
/*
 * fedi-imagebot: An imagebot for the Fediverse.
 * Copyright © 2021, Mick 🔥 Abernathy <@itsalltoast@to.ast.my>
 *   BSD-3 - See LICENSE for usage restrictions.
 */

import (
	"errors"
	"github.com/itsalltoast/fedi-imagebot/config"
	"github.com/mattn/go-mastodon"
)

// Mastodon stores configuration elements that are unique to the Mastodon API client.
//
type Mastodon struct {
	client    *mastodon.Client
	siteURL   string
	apiID     string
	apiSecret string
}

// newMastodon creates and prepares the Mastodon API client.
//
func newMastodon(cfg *config.Config) *Mastodon {
	m := new(Mastodon)
	m.siteURL = cfg.SiteURL
	m.apiID = cfg.SiteKey
	m.apiSecret = cfg.SiteSecret

	m.client = mastodon.NewClient(&mastodon.Config{
		Server:       m.siteURL,
		ClientID:     m.apiID,
		ClientSecret: m.apiSecret,
	})

	return m
}

// RemovePost (TODO)
//
func (m *Mastodon) RemovePost(postID string) error {
	return errors.New("Not implemented")
}

// PostImage (TODO)
func (m *Mastodon) PostImage(url string) error {
	return errors.New("Not implemented")
}
