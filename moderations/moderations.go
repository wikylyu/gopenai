package moderations

import (
	"encoding/json"
	"fmt"

	"github.com/wikylyu/gopenai/api"
)

func NewClient(c *api.Client) *ModerationClient {
	return &ModerationClient{c: c}
}

func (c *ModerationClient) Create(req *CreateRequest) (*CreateResponse, error) {
	switch req.Input.(type) {
	case string, []string:
		break
	default:
		return nil, fmt.Errorf("input must be string or array")
	}
	data, err := c.c.DoJson("POST", "/moderations", req)
	if err != nil {
		return nil, err
	}

	var resp CreateResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
