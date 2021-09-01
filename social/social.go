package social
/*
 * fedi-imagebot: An imagebot for the Fediverse.
 * Copyright © 2021, Mick 🔥 Abernathy <@itsalltoast@to.ast.my>
 *   BSD-3 - See LICENSE for usage restrictions.
 */

import (
	"github.com/itsalltoast/fedi-imagebot/config"
)

// Interface represents a base Social network connection object.
//
type Interface interface {
	RemovePost(postID string) error
	PostImage(url string) error
}

// Social stores the user's configuration object for this social API.
//
type Social struct {
	Interface
	config *config.Config
}

// NewSocial constructs a social network object for the user's requested
// backend API.  Currently only supports Misskey.
//
func NewSocial(c *config.Config) Interface {
	var s *Social
	switch c.SiteType {
	case "misskey":
		{
			sa := newMisskey(c)
			return sa
		}
	}
	return s
}
