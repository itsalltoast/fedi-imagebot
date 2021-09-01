/*
 * fedi-imagebot: An imagebot for the Fediverse.
 * Copyright © 2021, Mick 🔥 Abernathy <@itsalltoast@to.ast.my>
 *   BSD-3 - See LICENSE for usage restrictions.
 */
package social

import (
	"errors"
	"github.com/itsalltoast/fedi-imagebot/config"
	"github.com/mattn/go-mastodon"
)

type Mastodon struct {
	client    *mastodon.Client
	siteUrl   string
	apiID     string
	apiSecret string
}

func NewMastodon(cfg *config.Config) *Mastodon {
	m := new(Mastodon)
	m.siteUrl = cfg.SiteURL
	m.apiID = cfg.SiteKey
	m.apiSecret = cfg.SiteSecret

	m.client = mastodon.NewClient(&mastodon.Config{
		Server:       m.siteUrl,
		ClientID:     m.apiID,
		ClientSecret: m.apiSecret,
	})

	return m
}

func (m *Mastodon) RemovePost(postId string) error {
	return errors.New("Not implemented")
}

func (m *Mastodon) PostImage(url string) error {
	return errors.New("Not implemented")
}
