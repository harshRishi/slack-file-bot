package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error loading .env file: %s\n", err)
		os.Exit(1)
	}

	token := os.Getenv("SLACK_BOT_TOKEN")
	channelID := os.Getenv("CHANNEL_ID")
	filePath := "./test.pdf"

	client := slack.New(token)

	params := slack.FileUploadParameters{
		Channels:       []string{channelID},
		InitialComment: "Here's my file : Smile:",
		File:           filePath,
	}

	uploadedFile, err := client.UploadFile(params)
	if err != nil {
		fmt.Printf("Error uploading file: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Name: %s, URL: %s\n", uploadedFile.Name, uploadedFile.URL)
}
