package edits

import (
	"encoding/json"

	"github.com/wikylyu/gopenai/api"
)

func NewClient(c *api.Client) *EditClient {
	return &EditClient{c: c}
}

/*
 * Creates a new edit for the provided input, instruction, and parameters.
 */
func (c *EditClient) Create(req *CreateRequest) (*CreateResponse, error) {
	data, err := c.c.DoJson("POST", "/edits", req)
	if err != nil {
		return nil, err
	}

	var resp CreateResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
