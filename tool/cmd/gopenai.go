package cmd

import (
	"os"

	"github.com/wikylyu/gopenai"
)

var openai *gopenai.Session = nil

func init() {
	openai = gopenai.New(&gopenai.Config{
		ApiKey: os.Getenv("OPENAI_API_KEY"),
	})
}
