package search
/*
 * fedi-imagebot: An imagebot for the Fediverse.
 * Copyright © 2021, Mick 🔥 Abernathy <@itsalltoast@to.ast.my>
 *   BSD-3 - See LICENSE for usage restrictions.
 */

import (
	g "github.com/serpapi/google-search-results-golang"

	"crypto/rand"
	"errors"
	"math/big"

	"github.com/itsalltoast/fedi-imagebot/config"
)

// SerpAPI stores configuration variables specific to SerpAPI.
//
type SerpAPI struct {
	Search
	config *config.Config
}

func newSerpAPI(config *config.Config) *SerpAPI {
	s := new(SerpAPI)
	s.config = config

	return s
}

// GetURLSet returns a set of URLs matching the desired keywords.
//
func (s *SerpAPI) GetURLSet(keywords string) ([]string, error) {
	parameter := map[string]string{
		"engine":        "google",
		"q":             keywords,
		"google_domain": "google.com",
		"gl":            "us",
		"hl":            "en",
		"tbm":           "isch",
		"safe":          "active",
	}

	// Filter image aspect ratio at user request.  By default we pull ALL images.
	//
	if ImageAspect(s.config.AspectRatio) == ImageLandscape {
		parameter["tbs"] = "iar:w"
	} else if ImageAspect(s.config.AspectRatio) == ImagePortrait {
		parameter["tbs"] = "iar:t"
	} else if ImageAspect(s.config.AspectRatio) == ImageSquare {
		parameter["tbs"] = "iar:s"
	}

	// Select a random page from the search results.
	//
	page, err := rand.Int(rand.Reader, big.NewInt(30))
	if err != nil {
		return nil, err
	}
	parameter["ijn"] = string(page.Int64())

	query := g.NewGoogleSearch(parameter, s.config.SearchKey)
	rsp, err := query.GetJSON()
	if err != nil {
		return nil, errors.New("SerpAPI call failed:" + err.Error())
	}
	if results, ok := rsp["images_results"].([]interface{}); !ok {
		log.Println("search object doesn't contain images_results as expected")
	} else {
		resultSet := make([]string, 0)
		for _, result := range results {
			th := result.(map[string]interface{})
			resultSet = append(resultSet, th["original"].(string))
		}
		return resultSet, nil
	}

	return nil, errors.New("No results")
}

// GetRandomURL returns a single, random URL from a set of search results.
//
func (s *SerpAPI) GetRandomURL(keywords string) (string, error) {
	if res, err := s.GetURLSet(keywords); err == nil {
		var sel *big.Int
		sel, err = rand.Int(rand.Reader, big.NewInt(int64(len(res))))
		if err != nil {
			return "", err
		}

		return res[sel.Int64()], nil
	}
	return "", err
}
