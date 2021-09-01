module github.com/itsalltoast/fedi-imagebot/social

go 1.16

replace github.com/itsalltoast/fedi-imagebot/config v0.0.0 => ../config/

require (
	github.com/itsalltoast/fedi-imagebot/config v0.0.0
	github.com/mattn/go-mastodon v0.0.4
	github.com/yitsushi/go-misskey v1.0.1
)
