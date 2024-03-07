# GoMelody Discord Bot

A simple discord bot that queries YouTube for music and plays
it in a voice channel that the user that requested the music is in.

Written in Golang and relies on [discordgo](https://github.com/bwmarrin/discordgo).


### Self Hosting the Bot

To self-host the bot, you need to have three environment variables in the .env file.
The first two are the discord bot token and application id from 
the [Discord Developer Portal](https://discord.com/developers/applications) and
the final one is the YouTube API key from the [Google Developer Console](https://console.developers.google.com/).

Example .env file:

``` env
BOT_TOKEN=[string] # Discord bot token from the Discord Developer Portal
YOUTUBE_KEY=[string] 
APP_ID=[integer] # ID of the Discord application from the Discord Developer Portal
SERVER_ID=[integer] # ID of the Discord server the bot will be used in
```

To run the bot, simply install the dependencies and run the bot:

``` bash
go install
go run .
```
Or if you want to run it using Docker:
```
docker-compose up -d
```

### Available Commands

Currently, the bot supports the following commands:
 - `/play [query]` - Searches YouTube for the given query and plays the first result in the voice channel the user is in.
 - `/stop` - Stops the current song and leaves the voice channel.
 - `/leave` - Leaves the voice channel.
 - `/join` - Joins the voice channel the user is in.
 - `/ping` - Pings the bot and returns "Pong!" to test if the bot is online.


### Future Plans

The bot is more of a toy project, but if I have the time, I will probably add the following:
 - Support for more music sources other than YouTube
 - Queue system for playing music
 - Pausing and resuming songs
 - Support for shuffling the queue
 - Support for repeating songs