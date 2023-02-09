package embeddings

import "github.com/wikylyu/gopenai/client"

type CreateRequest struct {
	Model string      `json:"model"`
	Input interface{} `json:"input"` // string or []string
	User  string      `json:"user,omitempty"`
}

type Usage struct {
	PromptTokens int `json:"prompt_tokens"`
	TotalTokens  int `json:"total_tokens"`
}

type Embedding struct {
	Object    string    `json:"object"`
	Embedding []float64 `json:"embedding"`
	Index     int       `json:"index"`
}

type CreateResponse struct {
	Object string `json:"object"`
	Model  string `json:"model"`
	Usage  *Usage `json:"usage"`

	Data []*Embedding `json:"data"`
}

type EmbeddingClient struct {
	c *client.Client
}
