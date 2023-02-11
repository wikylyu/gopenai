package embeddings

import (
	"encoding/json"

	"github.com/wikylyu/gopenai/api"
	"github.com/wikylyu/gopenai/base"
)

func NewClient(c *api.Client) *EmbeddingClient {
	return &EmbeddingClient{c: c}
}

/*
 * Creates an embedding vector representing the input text.
 */
func (c *EmbeddingClient) Create(req *CreateRequest) (*CreateResponse, error) {
	if err := base.ValidateInput(req.Input); err != nil {
		return nil, err
	}
	data, err := c.c.DoJson("POST", "/embeddings", req)
	if err != nil {
		return nil, err
	}

	var resp CreateResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
