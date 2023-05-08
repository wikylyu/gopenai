package chat

import (
	"github.com/wikylyu/gopenai/api"
	"github.com/wikylyu/gopenai/base"
)

const (
	MessageRoleSystem    = "system"
	MessageRoleUser      = "user"
	MessageRoleAssistant = "assistant"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
	Name    string `json:"name"`
}

type CreateRequest struct {
	Model            string                 `json:"model"`
	Messages         []*Message             `json:"messages"`
	Temperature      float64                `json:"temperature,omitempty"`
	TopP             float64                `json:"top_p,omitempty"`
	N                int64                  `json:"n,omitempty"`
	Stream           bool                   `json:"stream,omitempty"`
	Stop             interface{}            `json:"stop,omitempty"` // string or array of string
	MaxTokens        int64                  `json:"max_tokens,omitempty"`
	PresencePenalty  float64                `json:"presence_penalty,omitempty"`
	FrequencyPenalty float64                `json:"frequency_penalty,omitempty"`
	LogitBias        map[string]interface{} `json:"logit_bias,omitempty"`
	User             string                 `json:"user,omitempty"`
}

type Choice struct {
	Index        int64    `json:"index"`
	Message      *Message `json:"message"`
	FinishReason string   `json:"stop"`
}

type CreateResponse struct {
	ID      string      `json:"id"`
	Object  string      `json:"object"`
	Created int64       `json:"created"`
	Choices []*Choice   `json:"choices"`
	Usage   *base.Usage `json:"usage"`
}

type ChatClient struct {
	c *api.Client
}
