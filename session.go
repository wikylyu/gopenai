package gopenai

import (
	"github.com/wikylyu/gopenai/client"
	"github.com/wikylyu/gopenai/completions"
	"github.com/wikylyu/gopenai/edits"
	"github.com/wikylyu/gopenai/models"
)

func New(cfg *Config) *Session {
	client := client.NewClient(cfg.Endpoint, cfg.ApiKey)
	return &Session{
		c:          client,
		Model:      models.NewModelClient(client),
		Completion: completions.NewCompletionClient(client),
		Edit:       edits.NewEditClient(client),
	}
}
