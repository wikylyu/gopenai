package edits

import "github.com/wikylyu/gopenai/client"

// https://platform.openai.com/docs/api-reference/edits/create
type CreateEditRequest struct {
	Model       string  `json:"model"`
	Input       string  `json:"input"`
	Instruction string  `json:"instruction"`
	N           int     `json:"n,omitempty"`
	Temperature float64 `json:"temperature,omitempty"`
	TopP        float64 `json:"top_p,omitempty"`
}

type CreateEditChoice struct {
	Text  string `json:"text"`
	Index int    `json:"index"`
}

type CreateEditUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type CreateEditResponse struct {
	Object  string              `json:"object"`
	Created int                 `json:"created"`
	Choices []*CreateEditChoice `json:"choices"`
	Usage   *CreateEditUsage    `json:"usage"`
}

type EditClient struct {
	c *client.Client
}
