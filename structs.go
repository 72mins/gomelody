package main

type VoiceChannel struct {
	ChannelID string
	GuildID   string
}

type Database struct {
	DpMode bool `json:"dp_mode"`
}
