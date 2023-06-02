package chat

import (
	"bufio"
	"net/http"
	"strings"

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
	Name    string `json:"name,omitempty"`
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

type StreamChoice struct {
	Delta struct {
		Content string `json:"content"`
	} `json:"delta"`
	Index        int64  `json:"index"`
	FinishReason string `json:"stop"`
}

type StreamResponse struct {
	ID      string          `json:"id"`
	Object  string          `json:"object"`
	Model   string          `json:"model"`
	Created int64           `json:"created"`
	Choices []*StreamChoice `json:"choices"`
}

func (r *StreamResponse) Content() string {
	var content string
	for _, choice := range r.Choices {
		if choice.Delta.Content != "" {
			content += choice.Delta.Content
		}
	}
	return content
}

type ChatStreamer struct {
	resp   *http.Response
	reader *bufio.Reader

	buf strings.Builder
}
