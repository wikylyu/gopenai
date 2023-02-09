package edits

import "github.com/wikylyu/gopenai/client"

// https://platform.openai.com/docs/api-reference/edits/create
type CreateRequest struct {
	Model       string  `json:"model"`
	Input       string  `json:"input"`
	Instruction string  `json:"instruction"`
	N           int64   `json:"n,omitempty"`
	Temperature float64 `json:"temperature,omitempty"`
	TopP        float64 `json:"top_p,omitempty"`
}

type Choice struct {
	Text  string `json:"text"`
	Index int64  `json:"index"`
}

type Usage struct {
	PromptTokens     int64 `json:"prompt_tokens"`
	CompletionTokens int64 `json:"completion_tokens"`
	TotalTokens      int64 `json:"total_tokens"`
}

type CreateResponse struct {
	Object  string    `json:"object"`
	Created int       `json:"created"`
	Choices []*Choice `json:"choices"`
	Usage   *Usage    `json:"usage"`
}

type EditClient struct {
	c *client.Client
}
