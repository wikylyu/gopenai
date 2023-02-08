package completions

import (
	"encoding/json"

	"github.com/wikylyu/gopenai/client"
)

func NewCompletionClient(c *client.Client) *CompletionClient {
	return &CompletionClient{c: c}
}

func (c *CompletionClient) Create(req *CreateCompletionRequest) (*CreateCompletionResponse, error) {
	data, err := c.c.Do("POST", "/completions", req)
	if err != nil {
		return nil, err
	}

	var resp CreateCompletionResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
