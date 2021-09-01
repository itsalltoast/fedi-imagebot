package main

import (
	"github.com/itsalltoast/fedi-imagebot/config"
	"os"
	"testing"
)

func TestValidEnv(t *testing.T) {
	os.Setenv("SITE_TYPE", "misskey")
	os.Setenv("SITE_URL", "https://nowhere")
	os.Setenv("SITE_KEY", "is set")
	os.Setenv("BOT_KEYWORDS", "umbrella in the rain")
	c := config.NewConfigFromEnv()
	if c == nil {
		t.Errorf("Config returned null pointer")
	}

	if !c.Valid() {
		t.Errorf("Config reported as invalid when it's fine.")
	}
}
