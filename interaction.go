package main

import (
	"fmt"
	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
	"os"
)

func InteractionResponse(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ApplicationCommandData()

	switch data.Name {
	case "ping":
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Pong!",
			},
		})
		if err != nil {
			return
		}
	case "join":
		channelInfo := FindVoiceChannel(s, i.Member.User.ID)

		if channelInfo.ChannelID == "" {
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "You are not in a voice channel!",
				},
			})
			if err != nil {
				return
			}
			return
		} else {
			_, err := s.ChannelVoiceJoin(channelInfo.GuildID, channelInfo.ChannelID, false, false)
			if err != nil {
				return
			}

			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Joined voice channel!",
				},
			})
			if err != nil {
				return
			}
		}
	case "leave":
		voice, err := s.ChannelVoiceJoin(i.GuildID, "", false, false)
		if err != nil {
			return
		}

		err = voice.Disconnect()
		if err != nil {
			return
		}

		err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Leaving voice channel :(",
			},
		})
	case "play":
		channelInfo := FindVoiceChannel(s, i.Member.User.ID)

		if channelInfo.ChannelID == "" {
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "You are not in a voice channel!",
				},
			})
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

			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Playing audio!",
				},
			})
			if err != nil {
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
