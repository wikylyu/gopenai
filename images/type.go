package images

import (
	"os"

	"github.com/wikylyu/gopenai/api"
)

type CreateRequest struct {
	Prompt         string `json:"prompt"`
	N              int64  `json:"n,omitempty"`
	Size           string `json:"size,omitempty"`
	ResponseFormat string `json:"response_format,omitempty"` // url or b64_json
	User           string `json:"user,omitempty"`
}

type ImagesResponse struct {
	Created int64 `json:"created"`
	Data    []struct {
		URL     string `json:"url"`
		B64Json string `json:"b64_json"`
	} `json:"data"`
}

type CreateEditRequest struct {
	Image          *os.File `json:"image" form:"image"`
	Mask           *os.File `json:"mask" form:"mask"`
	Prompt         string   `json:"prompt" form:"prompt"`
	N              int64    `json:"n,omitempty" form:"n"`
	Size           string   `json:"size,omitempty" form:"size"`
	ResponseFormat string   `json:"response_format,omitempty" form:"response_format"` // url or b64_json
	User           string   `json:"user,omitempty" form:"user"`
}

type CreateVariationRequest struct {
	Image          *os.File `json:"image" form:"image"`
	N              int64    `json:"n,omitempty" form:"n"`
	Size           string   `json:"size,omitempty" form:"size"`
	ResponseFormat string   `json:"response_format,omitempty" form:"response_format"` // url or b64_json
	User           string   `json:"user,omitempty" form:"user"`
}

type ImageClient struct {
	c *api.Client
}
