package models

import (
	"encoding/json"

	"github.com/wikylyu/gopenai/client"
)

func NewClient(c *client.Client) *ModelClient {
	return &ModelClient{c: c}
}

func (c *ModelClient) List() (*ListResponse, error) {
	data, err := c.c.DoJson("GET", "/models", nil)
	if err != nil {
		return nil, err
	}
	var resp ListResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *ModelClient) Retrieve(name string) (*Model, error) {
	data, err := c.c.DoJson("GET", "/models/"+name, nil)
	if err != nil {
		return nil, err
	}
	var resp Model
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
