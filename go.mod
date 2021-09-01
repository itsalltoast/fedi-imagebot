module github.com/itsalltoast/fedi-imagebot

go 1.16

replace github.com/itsalltoast/fedi-imagebot/search => ./search/

replace github.com/itsalltoast/fedi-imagebot/store => ./store/

replace github.com/itsalltoast/fedi-imagebot/social => ./social/

replace github.com/itsalltoast/fedi-imagebot/config => ./config/

require (
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/itsalltoast/fedi-imagebot/config v0.0.0
	github.com/itsalltoast/fedi-imagebot/search v0.0.0
	github.com/itsalltoast/fedi-imagebot/social v0.0.0
	github.com/itsalltoast/fedi-imagebot/store v0.0.0
	github.com/kr/text v0.2.0 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)
