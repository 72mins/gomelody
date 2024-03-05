package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	dl "github.com/kkdai/youtube/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"io"
	"os"
)

var (
	YoutubeKey string
)

func init() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// Get bot token from .env
	YoutubeKey = os.Getenv("YOUTUBE_KEY")
}

func YoutubeSearch(query string) string {
	ctx := context.Background()
	service, err := youtube.NewService(ctx, option.WithAPIKey(YoutubeKey))
	if err != nil {
		fmt.Println("Error creating YouTube service: ", err)
		return ""
	}

	call := service.Search.List([]string{"id", "snippet"}).Q(query).MaxResults(1)
	response, err := call.Do()
	if err != nil {
		fmt.Println("Error making search call: ", err)
		return ""
	}

	var (
		videoId, audioTitle string
	)

	for _, item := range response.Items {
		videoId = item.Id.VideoId
		audioTitle = item.Snippet.Title
	}

	if videoId == "" {
		// TODO: Handle no audio found
		fmt.Println("No audio found")
		return ""
	}

	ytClient := dl.Client{}

	vid, err := ytClient.GetVideo(videoId)
	if err != nil {
		fmt.Println("Error getting video: ", err)
		return ""
	}

	formats := vid.Formats.WithAudioChannels()
	stream, _, err := ytClient.GetStream(vid, &formats[0])
	if err != nil {
		fmt.Println("Error getting stream: ", err)
		return ""
	}

	defer func(stream io.ReadCloser) {
		err := stream.Close()
		if err != nil {
			fmt.Println("Error closing stream: ", err)
		}
	}(stream)

	file, err := os.Create("audio/" + audioTitle + ".mp3")
	if err != nil {
		fmt.Println("Error creating file: ", err)
		return ""
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file: ", err)
		}
	}(file)

	_, err = io.Copy(file, stream)
	if err != nil {
		fmt.Println("Error copying stream to file: ", err)
		return ""
	}

	fmt.Println("File created: ", audioTitle+".mp3")

	return audioTitle + ".mp3"
}
