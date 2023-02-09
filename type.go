package gopenai

import (
	"github.com/wikylyu/gopenai/completions"
	"github.com/wikylyu/gopenai/edits"
	"github.com/wikylyu/gopenai/embeddings"
	"github.com/wikylyu/gopenai/images"
	"github.com/wikylyu/gopenai/models"
)

type Session struct {
	// c          *client.Client
	Model      *models.ModelClient
	Completion *completions.CompletionClient
	Edit       *edits.EditClient
	Image      *images.ImagesClient
	Embeddings *embeddings.EmbeddingClient
}
