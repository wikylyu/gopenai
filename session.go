package gopenai

import (
	"github.com/wikylyu/gopenai/client"
	"github.com/wikylyu/gopenai/completions"
	"github.com/wikylyu/gopenai/edits"
	"github.com/wikylyu/gopenai/embeddings"
	"github.com/wikylyu/gopenai/images"
	"github.com/wikylyu/gopenai/models"
)

func New(cfg *Config) *Session {
	client := client.NewClient(cfg.Endpoint, cfg.ApiKey)
	return &Session{
		// c:          client,
		Model:      models.NewClient(client),
		Completion: completions.NewClient(client),
		Edit:       edits.NewClient(client),
		Image:      images.NewClient(client),
		Embeddings: embeddings.NewClient(client),
	}
}
