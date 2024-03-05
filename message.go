package main

import (
	"fmt"
	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
	"os"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate, i *discordgo.InteractionCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Get InteractionApplicationCommand

	if i != nil {
		fmt.Println("Interaction: ", i.Interaction.Message)
	}

	switch m.Content {
	case "!ping":
		_, err := s.ChannelMessageSend(m.ChannelID, "Pong!")
		if err != nil {
			return
		}
	case "!dl":
		YoutubeSearch()
	case "!join":
		channelInfo := FindVoiceChannel(s, m.Author.ID)

		if channelInfo.ChannelID == "" {
			_, err := s.ChannelMessageSend(m.ChannelID, "You are not in a voice channel!")
			if err != nil {
				return
			}
			return
		} else {
			_, err := s.ChannelVoiceJoin(channelInfo.GuildID, channelInfo.ChannelID, false, false)
			if err != nil {
				return
			}
		}
	case "!leave":
		voice, err := s.ChannelVoiceJoin(m.GuildID, "", false, false)
		if err != nil {
			return
		}

		err = voice.Disconnect()
	// Check if the message contains '!play' at the start
	case "!play":
		channelInfo := FindVoiceChannel(s, m.Author.ID)

		if channelInfo.ChannelID == "" {
			_, err := s.ChannelMessageSend(m.ChannelID, "You are not in a voice channel!")
			if err != nil {
				return
			}
			return
		} else {
			dgv, err := s.ChannelVoiceJoin(channelInfo.GuildID, channelInfo.ChannelID, false, false)
			if err != nil {
				fmt.Println("Error joining voice channel: ", err)
				return
			}

			YoutubeSearch()

			if !dgv.Ready {
				fmt.Println("Voice not ready")
				return
			}

			// get 'audio' folder from root
			Folder := "audio"

			// Start loop and attempt to play all files in the given folder
			fmt.Println("Reading Folder: ", Folder)
			files, _ := os.ReadDir(Folder)
			for _, f := range files {
				fmt.Println("PlayAudioFile:", f.Name())
				//discord.UpdateStatus(0, f.Name())

				dgvoice.PlayAudioFile(dgv, fmt.Sprintf("%s/%s", Folder, f.Name()), make(chan bool))
			}

			// Close connections
			dgv.Close()
			//discord.Close()

		}

	}

}
