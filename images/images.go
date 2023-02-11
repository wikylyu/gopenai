package images

import (
	"encoding/json"

	"github.com/wikylyu/gopenai/api"
	"github.com/wikylyu/gopenai/base"
)

func NewClient(c *api.Client) *ImageClient {
	return &ImageClient{c: c}
}

/*
 * Creates an image given a prompt.
 */
func (c *ImageClient) Create(req *CreateRequest) (*ImagesResponse, error) {
	if err := base.ValidateImageSize(req.Size); err != nil {
		return nil, err
	} else if err := base.ValidateImageFormat(req.ResponseFormat); err != nil {
		return nil, err
	}

	data, err := c.c.DoJson("POST", "/images/generations", req)
	if err != nil {
		return nil, err
	}

	var resp ImagesResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

/*
 * Creates an edited or extended image given an original image and a prompt.
 */
func (c *ImageClient) CreateEdit(req *CreateEditRequest) (*ImagesResponse, error) {
	if err := base.ValidateImageSize(req.Size); err != nil {
		return nil, err
	} else if err := base.ValidateImageFormat(req.ResponseFormat); err != nil {
		return nil, err
	}

	data, err := c.c.DoForm("POST", "/images/edits", req)
	if err != nil {
		return nil, err
	}

	var resp ImagesResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

/*
 * Creates a variation of a given image.
 */
func (c *ImageClient) CreateVariation(req *CreateVariationRequest) (*ImagesResponse, error) {
	if err := base.ValidateImageSize(req.Size); err != nil {
		return nil, err
	} else if err := base.ValidateImageFormat(req.ResponseFormat); err != nil {
		return nil, err
	}

	data, err := c.c.DoForm("POST", "/images/variations", req)
	if err != nil {
		return nil, err
	}

	var resp ImagesResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
