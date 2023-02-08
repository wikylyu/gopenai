package completions

import "github.com/wikylyu/gopenai/client"

type CreateCompletionRequest struct {
	Model       string  `json:"model"`
	Prompt      string  `json:"prompt"`
	MaxTokens   int     `json:"max_tokens,omitempty"`
	Temperature float64 `json:"temperature,omitempty"`
	TopP        float64 `json:"top_p,omitempty"`
	N           int     `json:"n,omitempty"`
	Stream      bool    `json:"stream,omitempty"`
	Logprobs    int     `json:"logprobs,omitempty"`
	Stop        string  `json:"stop,omitempty"`
}

type CompletionChoice struct {
	Text         string      `json:"text"`
	Index        int         `json:"index"`
	Logprobs     interface{} `json:"logprobs"`
	FinishReason string      `json:"finish_reason"`
}

type CompletionUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type CreateCompletionResponse struct {
	ID      string              `json:"id"`
	Object  string              `json:"object"`
	Created int                 `json:"created"`
	Model   string              `json:"model"`
	Choices []*CompletionChoice `json:"choices"`
	Usage   *CompletionUsage    `json:"usage"`
}

type CompletionClient struct {
	c *client.Client
}
