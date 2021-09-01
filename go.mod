module github.com/itsalltoast/fedi-imagebot

go 1.16

replace github.com/itsalltoast/fedi-imagebot/search => ./search/

replace github.com/itsalltoast/fedi-imagebot/store => ./store/

replace github.com/itsalltoast/fedi-imagebot/social => ./social/

replace github.com/itsalltoast/fedi-imagebot/config => ./config/

require (
	github.com/itsalltoast/fedi-imagebot/config v0.0.0
	github.com/itsalltoast/fedi-imagebot/search v0.0.0
	github.com/itsalltoast/fedi-imagebot/social v0.0.0
	github.com/itsalltoast/fedi-imagebot/store v0.0.0
	golang.org/x/lint v0.0.0-20210508222113-6edffad5e616 // indirect
	golang.org/x/tools v0.1.5 // indirect
)
