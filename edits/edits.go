package edits

import (
	"encoding/json"

	"github.com/wikylyu/gopenai/client"
)

func NewClient(c *client.Client) *EditClient {
	return &EditClient{c: c}
}

/*
 * Creates a new edit for the provided input, instruction, and parameters.
 */
func (c *EditClient) Create(req *CreateEditRequest) (*CreateEditResponse, error) {
	data, err := c.c.DoJson("POST", "/edits", req)
	if err != nil {
		return nil, err
	}

	var resp CreateEditResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
