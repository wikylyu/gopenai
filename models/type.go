package models

import "github.com/wikylyu/gopenai/client"

type Permission struct {
	ID                 string      `json:"id"`
	Object             string      `json:"object"`
	Created            int64       `json:"created"`
	AllowCreateEngine  bool        `json:"allow_create_engine"`
	AllowSampling      bool        `json:"allow_sampling"`
	AllowLogprobs      bool        `json:"allow_logprobs"`
	AllowSearchIndices bool        `json:"allow_search_indices"`
	AllowView          bool        `json:"allow_view"`
	AllowFineTuning    bool        `json:"allow_fine_tuning"`
	Organization       string      `json:"organization"`
	IsBlocking         bool        `json:"is_blocking"`
	Group              interface{} `json:"group"` // FIXME: don't know what it is.
}

type Model struct {
	ID         string        `json:"id"`
	Object     string        `json:"object"`
	OwnedBy    string        `json:"owned_by"`
	Created    int64         `json:"created"`
	Root       string        `json:"root"`
	Parent     interface{}   `json:"parent"` //FIXME: don't know what it is
	Permission []*Permission `json:"permission"`
}

type ListResponse struct {
	Data   []*Model `json:"data"`
	Object string   `json:"object"`
}

type ModelClient struct {
	c *client.Client
}
