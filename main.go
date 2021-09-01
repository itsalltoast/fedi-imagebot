package main
/*
 * fedi-imagebot: An imagebot for the Fediverse.
 * Copyright © 2021, Mick 🔥 Abernathy <@itsalltoast@to.ast.my>
 *   BSD-3 - See LICENSE for usage restrictions.
 */

import (
	"database/sql"
	"errors"
	"github.com/itsalltoast/fedi-imagebot/config"
	"github.com/itsalltoast/fedi-imagebot/search"
	"github.com/itsalltoast/fedi-imagebot/social"
	"github.com/itsalltoast/fedi-imagebot/store"
	"log"
	"os"
)

func discoverURLs(cfg *config.Config) []string {
	s := search.NewSearch(cfg)
	if r, e := s.GetURLSet(cfg.Keywords); e != nil {

		// API returned an error, we can't continue beyond this.
		//
		log.Fatal("Image search failed: ", e)

	} else {
		return r
	}

	return nil
}

func runAllRequestedBots(wantToGet bool, configs []*config.Config) {
	for _, cfg := range configs {
		db := store.NewStore(cfg)
		defer db.Close()

		// The user specifically only wants to add URLs.  This might be something they want to do weekly?
		//
		if wantToGet {
			r := discoverURLs(cfg)
			e := db.AddURLs(r)
			if e != nil {
				log.Println("Failed adding URLs:", e)
			}
			continue
		}

		url, e := db.GetRandomURL(true)

		// If we got this, we have no random URLs to select from.  Try and add some.
		//
		if e != nil && errors.Is(e, sql.ErrNoRows) {
			r := DiscoverURLs(cfg)
			e := db.AddURLs(r)
			if e != nil {
				log.Println(cfg.Name, e)
				continue
			}

			// If we were able to add some, hopefully this works this time.
			url, e = db.GetRandomURL(true)
		}

		// Note that this catches both possible errors above.
		if e != nil {
			log.Fatal(e)
		}

		mc := social.NewSocial(cfg)

		if e = mc.PostImage(url); e != nil {
			log.Println(cfg.Name, "Misskey API call failed: ", e)
			continue
		}
		// DEBUG: log.Println(">>", url)

		// Mark the URL as seen in the database.  Or die.
		//
		if e = db.MarkURL(url); e != nil {
			log.Println(cfg.Name, "Unable to mark URL as seen in the database:", e)
			continue
		}
	}
}

func main() {

	// Command line argument "get" instructs the bot to hit the search engine.  It will not do so
	// (by default) otherwise.
	//
	wantToGet := false

	// The only other arguments that can be passed in are JSON configuration files.
	//
	var configs []*config.Config
	for i, arg := range os.Args {

		// First argument is the command we're running.  Not interested.
		//
		if i < 1 {
			continue

			// The user wants us to find more content.
			//
		} else if arg == "get" {
			wantToGet = true
			continue
		}

		// Load the requested JSON file.
		//
		configFile, err := config.NewConfigFromFile(os.Args[1])
		if err != nil {
			log.Println("Failed loading configuration file:", arg)
			log.Println("Error:", err)
			continue

		} else if !configFile.Valid() {
			log.Println("Failed loading configuration file:", arg)
			log.Println("Ensure all required parameters are set.")
			continue
		}

		// If the JSON file hasn't named the bot, name the bot after the JSON file.
		//
		if len(configFile.Name) < 1 {
			configFile.Name = arg
		}

		// Add to the list of things to do.
		//
		configs = append(configs, configFile)
	}

	// If no arguments were passed in, we will try to load our configuration
	// from environment variables (intended for use with Docker).
	//
	if len(configs) < 1 {
		cfg := config.NewConfigFromEnv()
		if !cfg.Valid() {
			log.Fatalln("Unable to configure system from environment.")
		}
		configs = append(configs, cfg)
	}

	runAllRequestedBots(wantToGet, configs)
}
