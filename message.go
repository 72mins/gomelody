package main

import (
	"github.com/bwmarrin/discordgo"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	switch m.Content {
	case "!ping":
		_, err := s.ChannelMessageSend(m.ChannelID, "Pong!")
		if err != nil {
			return
		}
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
	case "!play":
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

			// TODO: Play music from youtube search
		}

	}

}
