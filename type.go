package gopenai

import (
	"github.com/wikylyu/gopenai/completions"
	"github.com/wikylyu/gopenai/edits"
	"github.com/wikylyu/gopenai/embeddings"
	"github.com/wikylyu/gopenai/files"
	"github.com/wikylyu/gopenai/finetunes"
	"github.com/wikylyu/gopenai/images"
	"github.com/wikylyu/gopenai/models"
	"github.com/wikylyu/gopenai/moderations"
)

type Session struct {
	Model      *models.ModelClient
	Completion *completions.CompletionClient
	Edit       *edits.EditClient
	Image      *images.ImageClient
	Embedding  *embeddings.EmbeddingClient
	File       *files.FileClient
	FineTune   *finetunes.FineTuneClient
	Moderation *moderations.ModerationClient
}
