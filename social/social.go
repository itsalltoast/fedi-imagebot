/*
 * fedi-imagebot: An imagebot for the Fediverse.
 * Copyright © 2021, Mick 🔥 Abernathy <@itsalltoast@to.ast.my>
 *   BSD-3 - See LICENSE for usage restrictions.
 */
package social

import (
	"github.com/itsalltoast/fedi-imagebot/config"
)

type SocialInterface interface {
	RemovePost(postId string) error
	PostImage(url string) error
}

type Social struct {
	SocialInterface
	config *config.Config
}

// Should we add base functionality or modularity here?  "Meh" for now
func NewSocial(c *config.Config) SocialInterface {
	var s *Social
	switch c.SiteType {
	case "misskey":
		{
			sa := NewMisskey(c)
			return sa
		}
	}
	return s
}
