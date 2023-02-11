package models

import (
	"encoding/json"

	"github.com/wikylyu/gopenai/api"
)

func NewClient(c *api.Client) *ModelClient {
	return &ModelClient{c: c}
}

/*
 * Lists the currently available models, and provides basic information about each one such as the owner and availability.
 */
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

/*
 * Retrieves a model instance, providing basic information about the model such as the owner and permissioning.
 * name: The ID of the model to use for this request
 */
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

/*
 * Delete a fine-tuned model. You must have the Owner role in your organization.
 */
func (c *ModelClient) Delete(id string) (*DeleteResponse, error) {
	data, err := c.c.DoJson("DELETE", "/models/"+id, nil)
	if err != nil {
		return nil, err
	}
	var resp DeleteResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
