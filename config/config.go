/*
 * fedi-imagebot: An imagebot for the Fediverse.
 * Copyright © 2021, Mick 🔥 Abernathy <@itsalltoast@to.ast.my>
 *   BSD-3 - See LICENSE for usage restrictions.
 */
package config

import (
	"os"
	"strconv"
	"io/ioutil"
	"encoding/json"
	"errors"
	"path/filepath"
)

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

var FileNotFound error = errors.New("File not found")
var FileIOError error = errors.New("Input/output error loading file")
var FileParseError error = errors.New("Parse error in config file")

func NewConfigFromEnv() *Config {
	config := new(Config)
	config.Name = os.Getenv("BOT_NAME")
	config.SiteType = os.Getenv("SITE_TYPE") // "misskey" / "mastodon" / "pleroma"
	config.SiteURL = os.Getenv("SITE_URL") // "https://xxxxx/"
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

func NewConfigFromFile(filename string) (*Config, error) {

	var config Config
	if f, e := os.Open(filepath.Clean(filename)); e != nil {
		return nil, FileNotFound
	} else {
		if data, e := ioutil.ReadAll(f); e != nil {
			return nil, FileIOError
		} else {
			e = json.Unmarshal(data, &config)
			if e != nil {
				return nil, FileParseError
			}

			return &config, nil
		}
	}
	return nil, errors.New("Unexpected error")
}

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
