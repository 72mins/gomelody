package main

import "github.com/bwmarrin/discordgo"

func GetApplicationCommands() []*discordgo.ApplicationCommand {
	return []*discordgo.ApplicationCommand{
		{
			Name:        "ping",
			Description: "Replies with Pong!",
		},
		{
			Name:        "join",
			Description: "Joins the voice channel you are in",
		},
		{
			Name:        "leave",
			Description: "Leaves the voice channel",
		},
		{
			Name:        "stop",
			Description: "Stops the current song",
		},
		{
			Name:        "play",
			Description: "Plays a song",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "name",
					Description: "Name of the song to play",
					Required:    true,
				},
			},
		},
	}
}
