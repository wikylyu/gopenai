package main

import (
	"fmt"
	"os"

	"github.com/wikylyu/gopenai"
	"github.com/wikylyu/gopenai/moderations"
)

func main() {
	openai := gopenai.New(&gopenai.Config{
		ApiKey: os.Getenv("OPENAI_API_KEY"),
	})
	resp, err := openai.Moderation.Create(&moderations.CreateRequest{
		Input: "I want to kill them!",
	})
	if err != nil {
		panic(err)
	}
	for _, r := range resp.Results {
		fmt.Printf("hate: %v\n", r.CategoryScores.Hate)
		fmt.Printf("Hate Threatening: %v\n", r.CategoryScores.HateThreatening)
		fmt.Printf("Self Harm: %v\n", r.CategoryScores.SelfHarm)
		fmt.Printf("Sexual: %v\n", r.CategoryScores.Sexual)
		fmt.Printf("Sexual Minors: %v\n", r.CategoryScores.SexualMinors)
		fmt.Printf("Violence : %v\n", r.CategoryScores.Violence)
		fmt.Printf("Violence Graphic: %v\n", r.CategoryScores.ViolenceGraphic)

	}
}
