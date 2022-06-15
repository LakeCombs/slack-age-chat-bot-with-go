package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {

	loadenv := godotenv.Load(".env")
	println("loadenv: ", loadenv)

	// os.Setenv("SLACK_BOT_TOKEN", "")
	// os.Setenv("CHANNEL_ID", "")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"cmp.pdf"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}

		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		fmt.Printf("name: %s, URL: %s\n", file.Name, file.URL)
	}

}
