package base

type ChoiceLogprobs struct {
	Tokens        []string      `json:"tokens"`
	TokenLogprobs []float64     `json:"token_logprobs"`
	TopLogprobs   []interface{} `json:"top_logprobs"`
	TextOffset    []int64       `json:"text_offset"`
}

type Choice struct {
	Text         string          `json:"text"`
	Index        int64           `json:"index"`
	Logprobs     *ChoiceLogprobs `json:"logprobs"`
	FinishReason string          `json:"finish_reason"`
}

type Usage struct {
	PromptTokens     int64 `json:"prompt_tokens"`
	CompletionTokens int64 `json:"completion_tokens"`
	TotalTokens      int64 `json:"total_tokens"`
}
