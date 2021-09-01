/*
 * fedi-imagebot: An imagebot for the Fediverse.
 * Copyright © 2021, Mick 🔥 Abernathy <@itsalltoast@to.ast.my>
 *   BSD-3 - See LICENSE for usage restrictions.
 */
package search

import (
	"errors"
	"github.com/itsalltoast/fedi-imagebot/config"
)

type SearchInterface interface {
	GetRandomURL(string) (string, error)
	GetURLSet(string) ([]string, error)
}

type Search struct {
	SearchAspect ImageAspect
	config       *config.Config
}

type ImageAspect int

const (
	ImageAny ImageAspect = iota
	ImageLandscape
	ImagePortrait
	ImageSquare
)

// Should we add base functionality or modularity here?  "Meh" for now
func NewSearch(c *config.Config /*searchType string, searchKey string, searchSecret string*/) SearchInterface {
	var s *Search
	switch c.SearchType {
	case "serpapi":
		{
			sa := NewSerpAPI(c)
			return sa
		}
	}
	return s
}

func (s *Search) SetAspect(aspect ImageAspect) {
	s.SearchAspect = aspect
}

func (s *Search) GetAspect() ImageAspect {
	return s.SearchAspect
}

func (s *Search) GetRandomURL(string) (string, error) {
	return "", errors.New("Cannot call base class")
}

func (s *Search) GetURLSet(string) ([]string, error) {
	return nil, errors.New("Cannot call base class")
}
