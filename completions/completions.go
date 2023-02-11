package completions

import (
	"encoding/json"

	"github.com/wikylyu/gopenai/api"
	"github.com/wikylyu/gopenai/base"
)

func NewClient(c *api.Client) *CompletionClient {
	return &CompletionClient{c: c}
}

/*
 * Creates a completion for the provided prompt and parameters
 */
func (c *CompletionClient) Create(req *CreateRequest) (*CreateResponse, error) {
	if err := base.ValidatePrompt(req.Prompt); err != nil {
		return nil, err
	}

	data, err := c.c.DoJson("POST", "/completions", req)
	if err != nil {
		return nil, err
	}

	var resp CreateResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
