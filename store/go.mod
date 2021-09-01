module github.com/itsalltoast/fedi-imagebot/store

go 1.16

replace github.com/itsalltoast/fedi-imagebot/config v0.0.0 => ../config/

require (
	github.com/itsalltoast/fedi-imagebot/config v0.0.0
	github.com/mattn/go-sqlite3 v2.0.3+incompatible
)
