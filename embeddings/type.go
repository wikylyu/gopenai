package embeddings

import (
	"github.com/wikylyu/gopenai/api"
	"github.com/wikylyu/gopenai/base"
)

type CreateRequest struct {
	Model string      `json:"model"`
	Input interface{} `json:"input"` // string or []string
	User  string      `json:"user,omitempty"`
}

type Embedding struct {
	Object    string    `json:"object"`
	Embedding []float64 `json:"embedding"`
	Index     int64     `json:"index"`
}

type CreateResponse struct {
	Object string      `json:"object"`
	Model  string      `json:"model"`
	Usage  *base.Usage `json:"usage"`

	Data []*Embedding `json:"data"`
}

type EmbeddingClient struct {
	c *api.Client
}
