package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"os"
	"os/signal"
)

var (
	ApplicationID string
)

func init() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// Get bot token from .env
	ApplicationID = os.Getenv("APP_ID")
}

func ConnectDiscord() {
	// Create a new Discord session using the provided bot token
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	GuildID := "539060061033463811"

	// Register commands
	_, err = dg.ApplicationCommandBulkOverwrite(ApplicationID, GuildID, []*discordgo.ApplicationCommand{
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
	})
	if err != nil {
		fmt.Println("Error creating slash commands: ", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events
	dg.AddHandler(InteractionResponse)

	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages | discordgo.IntentsGuilds | discordgo.IntentsGuildVoiceStates | discordgo.IntentsAll)

	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection: ", err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL+C to exit.")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	err = dg.Close()
	if err != nil {
		return
	}
}

func FindVoiceChannel(s *discordgo.Session, user string) VoiceChannel {
	for _, g := range s.State.Guilds {
		for _, v := range g.VoiceStates {
			if v.UserID == user {
				return VoiceChannel{v.ChannelID, g.ID}
			}
		}
	}

	return VoiceChannel{}
}
