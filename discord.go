package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

func ConnectDiscord() {
	// Create a new Discord session using the provided bot token
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events
	dg.AddHandler(MessageCreate)

	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages | discordgo.IntentsGuilds | discordgo.IntentsGuildVoiceStates | discordgo.IntentsAll)

	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection: ", err)
		return
	}

	_, err = dg.User("@me")
	if err != nil {
		fmt.Println("Error obtaining account details: ", err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL+C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

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
