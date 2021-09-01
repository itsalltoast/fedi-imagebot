package config

/*
 * fedi-imagebot: An imagebot for the Fediverse.
 * Copyright © 2021, Mick 🔥 Abernathy <@itsalltoast@to.ast.my>
 *   BSD-3 - See LICENSE for usage restrictions.
 */

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

// Config controls how the imagebot behaves.
//
type Config struct {
	Name            string `json:"name"`
	SiteType        string `json:"site"`
	SiteURL         string `json:"siteurl"`
	SiteKey         string `json:"sitekey"`
	SiteSecret      string `json:"sitesecret"`
	DBType          string `json:"db"`
	DBConnectString string `json:"dbconnect"`
	SearchType      string `json:"search"`
	SearchKey       string `json:"searchapi"`
	SearchSecret    string `json:"searchsecret"`
	Keywords        string `json:"keywords"`
	AspectRatio     int    `json:"aspect"`
	MisskeyFolder   string `json:"folder"`
	LowWaterMark    int    `json:"lwm"`
}

// ErrFileNotFound is returned when a configuration file cannot be found.
var ErrFileNotFound error = errors.New("File not found")

// ErrFileIOError is returned when a configuration file is found, but could not be read.
var ErrFileIOError error = errors.New("Input/output error loading file")

// ErrFileParseError is returned when a configuration file is not in correct JSON format.
var ErrFileParseError error = errors.New("Parse error in config file")

// NewConfigFromEnv returns a configuration object based on the environment variables set
// at imagebot runtime (intended for use in containers).
//
func NewConfigFromEnv() *Config {
	config := new(Config)
	config.Name = os.Getenv("BOT_NAME")
	config.SiteType = os.Getenv("SITE_TYPE") // "misskey" / "mastodon" / "pleroma"
	config.SiteURL = os.Getenv("SITE_URL")   // "https://xxxxx/"
	config.SiteKey = os.Getenv("SITE_KEY")
	config.SiteSecret = os.Getenv("SITE_SECRET")
	config.DBType = os.Getenv("DB_TYPE")
	config.DBConnectString = os.ExpandEnv(os.Getenv("DB_CONNECT"))
	config.SearchType = os.Getenv("SEARCH_TYPE")
	config.SearchKey = os.Getenv("SEARCH_KEY")
	config.Keywords = os.Getenv("BOT_KEYWORDS")
	t, _ := strconv.Atoi(os.Getenv("ASPECT_RATIO"))
	config.MisskeyFolder = os.Getenv("MISSKEY_DRIVE_FOLDERID")
	config.AspectRatio = t
	return config
}

// NewConfigFromFile returns a configuration object based on a JSON file.
//
func NewConfigFromFile(filename string) (*Config, error) {
	var config Config
	if f, e := os.Open(filepath.Clean(filename)); e == nil {
		var data []byte
		var e error
		if data, e = ioutil.ReadAll(f); e == nil {
			e = json.Unmarshal(data, &config)
			if e != nil {
				return nil, ErrFileParseError
			}

			return &config, nil
		}

		if e != nil {
			return nil, ErrFileIOError
		}
	}

	return nil, ErrFileNotFound
}

// Valid performs validation of a config object (and in minimal cases, inserts defaults
// where appropriate).  Returns false on invalid configuration.
//
func (c *Config) Valid() bool {
	// DATABASE: By default, we will use a SQLite3 database in the user's home directory.
	//
	if len(c.DBType) < 1 {
		c.DBType = "sqlite3"
	}

	// Without these values, we cannot continue.
	//
	if len(c.SiteType) < 1 || len(c.SiteURL) < 1 || len(c.SiteKey) < 1 || len(c.Keywords) < 1 {
		return false
	}

	// Everything else should be fine.
	//
	return true
}
