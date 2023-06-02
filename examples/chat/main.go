package main

import (
	"fmt"
	"os"

	"github.com/wikylyu/gopenai"
	"github.com/wikylyu/gopenai/api"
	"github.com/wikylyu/gopenai/chat"
	"github.com/wikylyu/gopenai/models"
)

func main() {
	openai := gopenai.New(&gopenai.Config{
		ApiKey:  os.Getenv("OPENAI_API_KEY"),
		BaseURL: "http://43.154.167.208/v1",
	})

	streamer, err := openai.Chat.CreateStream(&chat.CreateRequest{
		Model:       models.ModelGPT3_5_Turbo,
		Temperature: 0.5,
		Messages: []*chat.Message{
			{
				Role:    chat.MessageRoleUser,
				Content: "send a http request with python3",
			},
		},
	})
	if err != nil {
		panic(err)
	}
	defer streamer.Close()

	for {
		resp, err := streamer.Read()
		if err != nil {
			if apierr := err.(*api.Error); apierr != nil {
				fmt.Printf("API Error: %v %v", apierr.Code, apierr.Message)
				break
			}
			panic(err)
		} else if resp == nil { // EOF
			break
		}
		for _, choice := range resp.Choices {
			if choice.Delta.Content != "" {
				fmt.Printf("%s", choice.Delta.Content)
			}
		}
	}
	fmt.Printf("\n=============================================\n")

	fmt.Printf("%s\n", streamer.Content())
}
