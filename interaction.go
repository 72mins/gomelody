package main

import (
	"fmt"
	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
	"os"
	"path/filepath"
)

var (
	dgv *discordgo.VoiceConnection
)

func CleanAudioFolder() error {
	path := "audio"

	d, err := os.Open(path)
	if err != nil {
		return err
	}

	defer func(d *os.File) {
		err := d.Close()
		if err != nil {
			fmt.Println("Error closing file: ", err)
		}
	}(d)

	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}

	// Remove all files in the audio folder
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(path, name))
		if err != nil {
			return err
		}
	}

	return nil
}

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
	case "leave", "stop":
		voice, err := s.ChannelVoiceJoin(i.GuildID, "", false, false)
		if err != nil {
			return
		}

		err = voice.Disconnect()
		if err != nil {

			return
		}

		err = CleanAudioFolder()
		if err != nil {
			return
		}

		err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Stopped audio and left voice channel.",
			},
		})
	case "play":
		channelInfo := FindVoiceChannel(s, i.Member.User.ID)

		files, _ := os.ReadDir("audio")
		if len(files) > 0 {
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "A song is already playing!",
				},
			})
			if err != nil {
				return
			}
			return
		}

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
			if dgv == nil {
				dgv, _ = s.ChannelVoiceJoin(channelInfo.GuildID, channelInfo.ChannelID, false, false)
			} else {
				dgv.Close()

				dgv, _ = s.ChannelVoiceJoin(channelInfo.GuildID, channelInfo.ChannelID, false, false)
			}

			// Get query from the user and search for the song on YouTube
			query := data.Options[0].Value.(string)
			title := YoutubeSearch(i, s, query)

			_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Playing " + title,
				},
			})

			if !dgv.Ready {
				fmt.Println("Voice not ready")
				return
			}

			Folder := "audio"

			// Play all files in the audio folder
			files, _ := os.ReadDir(Folder)
			for _, f := range files {
				dgvoice.PlayAudioFile(dgv, fmt.Sprintf("%s/%s", Folder, f.Name()), make(chan bool))
			}

			err := CleanAudioFolder()
			if err != nil {
				return
			}

			err = dgv.Disconnect()
			if err != nil {
				return
			}
		}
	}
}
