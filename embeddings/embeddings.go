package embeddings

import (
	"encoding/json"
	"fmt"

	"github.com/wikylyu/gopenai/client"
)

func NewClient(c *client.Client) *EmbeddingClient {
	return &EmbeddingClient{c: c}
}

func (c *EmbeddingClient) Create(req *CreateRequest) (*CreateResponse, error) {
	switch req.Input.(type) {
	case string, []string:
		break
	default:
		return nil, fmt.Errorf("input should be string or array")
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
