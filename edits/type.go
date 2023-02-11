package edits

import (
	"github.com/wikylyu/gopenai/api"
	"github.com/wikylyu/gopenai/base"
)

// https://platform.openai.com/docs/api-reference/edits/create
type CreateRequest struct {
	Model       string  `json:"model"`
	Input       string  `json:"input,omitempty"`
	Instruction string  `json:"instruction"`
	N           int64   `json:"n,omitempty"`
	Temperature float64 `json:"temperature,omitempty"`
	TopP        float64 `json:"top_p,omitempty"`
}

type CreateResponse struct {
	Object  string         `json:"object"`
	Created int            `json:"created"`
	Choices []*base.Choice `json:"choices"`
	Usage   *base.Usage    `json:"usage"`
}

type EditClient struct {
	c *api.Client
}
