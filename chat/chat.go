package chat

import (
	"encoding/json"

	"github.com/wikylyu/gopenai/api"
)

func NewClient(c *api.Client) *ChatClient {
	return &ChatClient{c: c}
}

func (c *ChatClient) Create(req *CreateRequest) (*CreateResponse, error) {

	data, err := c.c.DoJson("POST", "/chat/completions", req)
	if err != nil {
		return nil, err
	}

	var resp CreateResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
