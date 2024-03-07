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
	ServerID      string
)

func init() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	ApplicationID = os.Getenv("APP_ID")
	ServerID = os.Getenv("SERVER_ID")
}

func ConnectDiscord() {
	// Create a new Discord session using the provided bot token
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	// Register commands
	_, err = dg.ApplicationCommandBulkOverwrite(ApplicationID, ServerID, GetApplicationCommands())
	if err != nil {
		fmt.Println("Error creating slash commands: ", err)
		return
	}

	// Register the InteractionResponse func as a callback for MessageCreate events
	dg.AddHandler(InteractionResponse)

	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages | discordgo.IntentsGuilds | discordgo.IntentsGuildVoiceStates)

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
