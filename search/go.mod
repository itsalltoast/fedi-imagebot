module github.com/itsalltoast/fedi-imagebot/search

go 1.16

replace github.com/itsalltoast/fedi-imagebot/config => ../config/

require (
	github.com/itsalltoast/fedi-imagebot/config v0.0.0
	github.com/serpapi/google-search-results-golang v0.0.0-20200815030216-632c97dac1ab
)
