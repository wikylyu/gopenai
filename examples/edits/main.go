package main

import (
	"fmt"
	"os"

	"github.com/wikylyu/gopenai"
	"github.com/wikylyu/gopenai/edits"
	"github.com/wikylyu/gopenai/models"
)

func main() {
	openai := gopenai.New(&gopenai.Config{
		ApiKey: os.Getenv("OPENAI_API_KEY"),
	})

	resp, err := openai.Edit.Create(&edits.CreateRequest{
		Model:       models.ModelTextDavinciEdit001,
		Input:       "What day of the wek is it?",
		Instruction: "Fix the spelling mistakes",
	})
	if err != nil {
		panic(err)
	}
	for _, choice := range resp.Choices {
		fmt.Printf("%s\n", choice.Text)
	}
}
