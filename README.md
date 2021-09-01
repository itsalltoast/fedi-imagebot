# Misskey Imagebot

A very basic bot that finds random images (currently using SerpAPI) and posts them to a Misskey timeline.

## How to use

# Compile/build:
There's nothing special about compiling this:

```
git clone https://github.com/itsalltoast/go-misskey-imagebot.git
cd go-misskey-imagebot
go build
```

# Configuration:

The app is currently reading all of its configuration from environment variables.  I'll be working on adding JSON config file and/or commandline support.

| Environment Variable | Use |
|----------------------|-----|
| SEARCH_TYPE | serpapi / google / bing / yandex (only serpapi currently supported) |
| SEARCH_KEY | An API key used to access your search API |
| BOT_KEYWORDS | Keywords to use in the image search |
| SITE_TYPE | misskey / pleroma / mastodon (only misskey currently supported) |
| SITE_URL | URL to your ActivityPub instance |
| SITE_KEY | API key for your target user account |
| MISSKEY_DRIVE_FOLDERID | (optional) Uploads the image to the account's specified Drive folder (only supported on Misskey instances) |
| BOT_NAME | (optional) backend name for the bot, currently only used in log messages |
| DB_TYPE | sqlite3 / pgsql / mysql / mongo (only sqlite3 currently supported) |
| DB_CONNECT | Connect string for the database instance |
| ASPECT_RATIO | (optional) 1=Landscape only, 2=Portrait only, 3=Square only.  0 or Unset = all images)


# Data Storage:

Runtime data for the bot (discovered images, post timestamps) is currently stored in `${HOME}/.go-misskey.imagebot.db` (sqlite3).

# Deployment:

## Run via Cron
1. Create a json file containing your configuration, such as `${HOME}/.MyNiceRobot.json`.
2. Run `crontab -e`
3. Make your robot run hourly:

```
0 * * * * ~/.bin/fedi-imagebot ~/.MyNiceRobot.json
```

4. Enjoy the content!

## Run via Docker
1. _I need to build a container image for this... Soon!_ :)
2. Use `-e` to configure the container instance for your robot.
?. _I need to add support for internal post scheduling (instead of external scheduling via cron)._
4. Enjoy the content!

```
$ docker run -it --rm --name mynicebot \
   -e SEARCH_KEY="xxxx" -e BOT_KEYWORDS="lain" \
   -e SITE_URL="https://me.social" \
   -e SITE_KEY="xxxx" \
      itsalltoast/fedi-imagebot
```

# Limitations:
Only supports Misskey and SerpAPI at present.

## Coming Soon

I'm working on making this better:

:white_check_mark: Configuration via json file and/or commandline parameters, instead of environment variables.

:white_check_mark: Support for multiple bot "personalities" (user account/keyword combinations).

:x: Support for other search APIs beyond SerpAPI.

:x: Support for other database backends beyond sqlite3 (Planned: mysql, postgresql, mongodb).

:x: Support for Mastodon (in progress).

:x: Support for Pleroma (requires a go-pleroma library, which I'll hopefully start soon).

:x: The ability to provide strings for the bot to include with the image as the post text.

:x: KNOWN ISSUE: If the bot runs across an image that is hotlink-protected, it will faithfully upload the html as an image and try to post it.  Need to add a HTTP HEAD call to the image ahead of time to ensure it is the expected mime-type.

:x: Run scheduling internally via a timer, to allow us to run this inside a standalone container

:x: Lambda support (why not?)
