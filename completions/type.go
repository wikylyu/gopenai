package completions

import (
	"github.com/wikylyu/gopenai/api"
	"github.com/wikylyu/gopenai/base"
)

type CreateRequest struct {
	Model            string                 `json:"model"`
	Prompt           interface{}            `json:"prompt,omitempty"` // string, array of string, array of int64, array of array of int64.
	Suffix           string                 `json:"suffix,omitempty"`
	MaxTokens        int64                  `json:"max_tokens,omitempty"`
	Temperature      float64                `json:"temperature,omitempty"`
	TopP             float64                `json:"top_p,omitempty"`
	N                int64                  `json:"n,omitempty"`
	Stream           bool                   `json:"stream,omitempty"`
	Logprobs         int64                  `json:"logprobs,omitempty"`
	Echo             bool                   `json:"echo,omitempty"`
	Stop             interface{}            `json:"stop,omitempty"` // string or array of string
	PresencePenalty  float64                `json:"presence_penalty,omitempty"`
	FrequencyPenalty float64                `json:"frequency_penalty,omitempty"`
	BestOf           int64                  `json:"best_of,omitempty"`
	LogitBias        map[string]interface{} `json:"logit_bias,omitempty"`
	User             string                 `json:"user,omitempty"`
}

type CreateResponse struct {
	ID      string         `json:"id"`
	Object  string         `json:"object"`
	Created int64          `json:"created"`
	Model   string         `json:"model"`
	Choices []*base.Choice `json:"choices"`
	Usage   *base.Usage    `json:"usage"`
}

type CompletionClient struct {
	c *api.Client
}
