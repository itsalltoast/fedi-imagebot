package search
/*
 * fedi-imagebot: An imagebot for the Fediverse.
 * Copyright © 2021, Mick 🔥 Abernathy <@itsalltoast@to.ast.my>
 *   BSD-3 - See LICENSE for usage restrictions.
 */

import (
	"errors"
	"github.com/itsalltoast/fedi-imagebot/config"
)

// Interface provides a common set of functions useful 
// with various search engines.
//
type Interface interface {
	GetRandomURL(string) (string, error)
	GetURLSet(string) ([]string, error)
}

// Search provides a link back to the user's selected
// configuration object.
//
type Search struct {
	SearchAspect ImageAspect
	config       *config.Config
}

// ImageAspect is a marker for aspect ratio preferences that
// will be translated and passed into the backend search
// engines (when possible).
//
type ImageAspect int

const (
	// ImageAny By default all images are acceptable.
	ImageAny ImageAspect = iota

	// ImageLandscape Prefer landscape images.
	ImageLandscape

	// ImagePortrait Prefer portrait images.
	ImagePortrait

	// ImageSquare Prefer square images.
	ImageSquare
)

// NewSearch prepares and returns the appropriate search object as requested
// by the user's configuration object.
//
func NewSearch(c *config.Config) Interface {
	var s *Search
	switch c.SearchType {
	case "serpapi":
		{
			sa := newSerpAPI(c)
			return sa
		}
	}
	return s
}

// SetAspect updates the aspect ratio parameter in this search object.
// TODO: unnecessary?
//
func (s *Search) SetAspect(aspect ImageAspect) {
	s.SearchAspect = aspect
}

// GetAspect returns the aspect ratio parameter from this search object.
// TODO: unnecessary?
//
func (s *Search) GetAspect() ImageAspect {
	return s.SearchAspect
}

// GetRandomURL is probably unnecessary (TODO)
//
func (s *Search) GetRandomURL(string) (string, error) {
	return "", errors.New("Cannot call base class")
}

// GetURLSet is probably unnecessary (TODO)
//
func (s *Search) GetURLSet(string) ([]string, error) {
	return nil, errors.New("Cannot call base class")
}
