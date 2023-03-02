package audio

import (
	"os"

	"github.com/wikylyu/gopenai/api"
)

type TranscribeRequest struct {
	File           *os.File `json:"file" form:"file"`
	Model          string   `json:"model" form:"model"`
	Prompt         string   `json:"prompt,omitempty" form:"prompt"`
	ResponseFormat string   `json:"response_format" form:"response_format"`
	Temperature    float64  `json:"temperature" form:"temperature"`
	Language       string   `json:"language" form:"language"`
}

type TranscribeResponse struct {
	Text string `json:"text"`
}

type TranslateRequest struct {
	File           *os.File `json:"file" form:"file"`
	Model          string   `json:"model" form:"model"`
	Prompt         string   `json:"prompt,omitempty" form:"prompt"`
	ResponseFormat string   `json:"response_format" form:"response_format"`
	Temperature    float64  `json:"temperature" form:"temperature"`
}

type TranslateResponse struct {
	Text string `json:"text"`
}

type AudioClient struct {
	c *api.Client
}
