package completions

import "github.com/wikylyu/gopenai/client"

type CreateRequest struct {
	Model       string  `json:"model"`
	Prompt      string  `json:"prompt"`
	MaxTokens   int64   `json:"max_tokens,omitempty"`
	Temperature float64 `json:"temperature,omitempty"`
	TopP        float64 `json:"top_p,omitempty"`
	N           int64   `json:"n,omitempty"`
	Stream      bool    `json:"stream,omitempty"`
	Logprobs    int64   `json:"logprobs,omitempty"`
	Stop        string  `json:"stop,omitempty"`
}

type Choice struct {
	Text         string      `json:"text"`
	Index        int64       `json:"index"`
	Logprobs     interface{} `json:"logprobs"`
	FinishReason string      `json:"finish_reason"`
}

type Usage struct {
	PromptTokens     int64 `json:"prompt_tokens"`
	CompletionTokens int64 `json:"completion_tokens"`
	TotalTokens      int64 `json:"total_tokens"`
}

type CreateResponse struct {
	ID      string    `json:"id"`
	Object  string    `json:"object"`
	Created int64     `json:"created"`
	Model   string    `json:"model"`
	Choices []*Choice `json:"choices"`
	Usage   *Usage    `json:"usage"`
}

type CompletionClient struct {
	c *client.Client
}
