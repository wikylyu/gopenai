package gopenai

import (
	"github.com/wikylyu/gopenai/api"
	"github.com/wikylyu/gopenai/chat"
	"github.com/wikylyu/gopenai/completions"
	"github.com/wikylyu/gopenai/edits"
	"github.com/wikylyu/gopenai/embeddings"
	"github.com/wikylyu/gopenai/files"
	"github.com/wikylyu/gopenai/finetunes"
	"github.com/wikylyu/gopenai/images"
	"github.com/wikylyu/gopenai/models"
	"github.com/wikylyu/gopenai/moderations"
)

func New(cfg *Config) *Session {
	client := api.NewClient(cfg.BaseURL, cfg.ApiKey, cfg.MaxIdleConns, cfg.IdleConnTimeout)
	return &Session{
		Model:      models.NewClient(client),
		Completion: completions.NewClient(client),
		Edit:       edits.NewClient(client),
		Image:      images.NewClient(client),
		Embedding:  embeddings.NewClient(client),
		File:       files.NewClient(client),
		FineTune:   finetunes.NewClient(client),
		Moderation: moderations.NewClient(client),
		Chat:       chat.NewClient(client),
	}
}
