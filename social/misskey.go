/*
 * fedi-imagebot: An imagebot for the Fediverse.
 * Copyright © 2021, Mick 🔥 Abernathy <@itsalltoast@to.ast.my>
 *   BSD-3 - See LICENSE for usage restrictions.
 */
package social

import (
	"github.com/yitsushi/go-misskey"
	//"github.com/yitsushi/go-misskey/services/drive"
	"github.com/itsalltoast/fedi-imagebot/config"
	"github.com/yitsushi/go-misskey/services/drive/files"
	"github.com/yitsushi/go-misskey/services/notes"
	"os"
)

type Misskey struct {
	client   *misskey.Client
	siteUrl  string
	apiKey   string
	folderId string
}

func NewMisskey(cfg *config.Config) *Misskey {
	m := new(Misskey)
	m.siteUrl = cfg.SiteURL
	m.apiKey = cfg.SiteKey
	m.folderId = cfg.MisskeyFolder
	m.client = misskey.NewClient(m.siteUrl, m.apiKey)

	return m
}

func (m *Misskey) RemovePost(postId string) error {
	postToRemove := os.Args[1]

	if note, err := m.client.Notes().Show(postToRemove); err != nil {
		return err
	} else {
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
	}
	return nil
}

func (m *Misskey) PostImage(url string) error {

	//	name := RandomString(24)
	if file, err := m.client.Drive().File().CreateFromURL(files.CreateFromURLOptions{
		FolderID: m.folderId,
		//		Name:     name,
		URL: url,
	}); err != nil {
		return err
	} else {
		if _, err := m.client.Notes().Create(notes.CreateRequest{
			Visibility: "public",
			FileIDs:    []string{file.ID},
		}); err != nil {
			return err
		}
	}
	return nil
}
