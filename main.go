package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

var (
	Token string
)

func init() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// Get bot token from .env
	Token = os.Getenv("BOT_TOKEN")
}

func main() {
	// Connect to Discord
	ConnectDiscord()
}
