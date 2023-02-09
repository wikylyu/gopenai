package images

import (
	"encoding/json"

	"github.com/wikylyu/gopenai/client"
)

func NewClient(c *client.Client) *ImageClient {
	return &ImageClient{c: c}
}

/*
 * Creates an image given a prompt.
 */
func (c *ImageClient) Create(req *CreateRequest) (*CreateResponse, error) {
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

/*
 * Creates an edited or extended image given an original image and a prompt.
 */
func (c *ImageClient) CreateEdit(req *CreateEditRequest) (*CreateEditResponse, error) {
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

/*
 * Creates a variation of a given image.
 */
func (c *ImageClient) CreateVariation(req *CreateVariationRequest) (*CreateVariationResponse, error) {
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
