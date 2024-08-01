package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	sklId := os.Getenv("SLACK_BOT_TOKEN")
	chId := os.Getenv("CHANNEL_ID")

	fmt.Println("My Creds: ", sklId, chId)

	// api := slack.New()
	// channelArr := []string{}
	// fileArr := []string{""}
}
