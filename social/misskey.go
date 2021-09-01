package social
/*
 * fedi-imagebot: An imagebot for the Fediverse.
 * Copyright © 2021, Mick 🔥 Abernathy <@itsalltoast@to.ast.my>
 *   BSD-3 - See LICENSE for usage restrictions.
 */

import (
	"github.com/yitsushi/go-misskey"
	//"github.com/yitsushi/go-misskey/services/drive"
	"github.com/itsalltoast/fedi-imagebot/config"
	"github.com/yitsushi/go-misskey/services/drive/files"
	"github.com/yitsushi/go-misskey/services/notes"
	"os"
)

// Misskey stores settings that are specific to using the Misskey API.
//
type Misskey struct {
	client   *misskey.Client
	siteURL  string
	apiKey   string
	folderID string
}

func newMisskey(cfg *config.Config) *Misskey {
	m := new(Misskey)
	m.siteURL = cfg.SiteURL
	m.apiKey = cfg.SiteKey
	m.folderID = cfg.MisskeyFolder
	m.client = misskey.NewClient(m.siteURL, m.apiKey)

	return m
}

// RemovePost will delete the identified post and its associated file.
//
func (m *Misskey) RemovePost(postID string) error {
	postToRemove := os.Args[1]

	if note, err := m.client.Notes().Show(postToRemove); err == nil {
		var fileToRemove string
		if len(note.FileIds) > 0 {
			fileToRemove = note.FileIds[0]
		}

		if err = m.client.Notes().Delete(postToRemove); err != nil {
			return err
		}

		if err = m.client.Drive().File().Delete(fileToRemove); err != nil {
			return err
		}
	} else {
		return err
	}
	return nil
}

// PostImage will upload the image to the user's Drive and post the image to the timeline.
//
func (m *Misskey) PostImage(url string) error {

	//	name := RandomString(24)
	if file, err := m.client.Drive().File().CreateFromURL(files.CreateFromURLOptions{
		FolderID: m.folderID,
		//		Name:     name,
		URL: url,
	}); err == nil {
		if _, err := m.client.Notes().Create(notes.CreateRequest{
			Visibility: "public",
			FileIDs:    []string{file.ID},
		}); err != nil {
			return err
		}
	} else {
		return err
	}
	return nil
}
