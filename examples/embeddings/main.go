package main

import (
	"fmt"
	"os"

	"github.com/wikylyu/gopenai"
	"github.com/wikylyu/gopenai/embeddings"
	"github.com/wikylyu/gopenai/models"
)

func main() {
	openai := gopenai.New(&gopenai.Config{
		ApiKey: os.Getenv("OPENAI_API_KEY"),
	})

	resp, err := openai.Embeddings.Create(&embeddings.CreateRequest{
		Model: models.ModelTextEmbeddingAda002,
		Input: "The food was delicious and the waiter...", // input can be string or array
	})
	if err != nil {
		panic(err)
	}
	for _, embedding := range resp.Data {
		fmt.Printf("%v\n", embedding.Embedding)
	}
}
