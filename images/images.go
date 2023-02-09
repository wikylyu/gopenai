package images

import (
	"encoding/json"

	"github.com/wikylyu/gopenai/client"
)

func NewClient(c *client.Client) *ImagesClient {
	return &ImagesClient{c: c}
}

func (c *ImagesClient) Create(req *CreateRequest) (*CreateResponse, error) {
	data, err := c.c.DoJson("POST", "/images/generations", req)
	if err != nil {
		return nil, err
	}

	var resp CreateResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *ImagesClient) CreateEdit(req *CreateEditRequest) (*CreateEditResponse, error) {
	data, err := c.c.DoForm("POST", "/images/edits", req)
	if err != nil {
		return nil, err
	}

	var resp CreateEditResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *ImagesClient) CreateVariation(req *CreateVariationRequest) (*CreateVariationResponse, error) {
	data, err := c.c.DoForm("POST", "/images/variations", req)
	if err != nil {
		return nil, err
	}

	var resp CreateVariationResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
