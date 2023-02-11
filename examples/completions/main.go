package main

import (
	"fmt"
	"os"

	"github.com/wikylyu/gopenai"
	"github.com/wikylyu/gopenai/completions"
	"github.com/wikylyu/gopenai/models"
)

func main() {
	openai := gopenai.New(&gopenai.Config{
		ApiKey: os.Getenv("OPENAI_API_KEY"),
	})

	prompt := []string{"Say this is a test"}
	resp, err := openai.Completion.Create(&completions.CreateRequest{
		Model:       models.ModelTextDavinci003,
		Prompt:      prompt,
		MaxTokens:   256,
		Temperature: 0,
	})
	if err != nil {
		panic(err)
	}
	for _, choice := range resp.Choices {
		fmt.Printf("%s\n", choice.Text)
	}
}
