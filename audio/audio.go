package audio

import (
	"encoding/json"

	"github.com/wikylyu/gopenai/api"
)

func NewClient(c *api.Client) *AudioClient {
	return &AudioClient{c: c}
}

/*
 * Transcribes audio into the input language.
 */
func (c *AudioClient) Transcribe(req *TranscribeRequest) (*TranscribeResponse, error) {
	data, err := c.c.DoForm("POST", "/audio/transcriptions", req)
	if err != nil {
		return nil, err
	}
	var resp TranscribeResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

/*
 * Translates audio into into English.
 */
func (c *AudioClient) Translate(req *TranslateRequest) (*TranslateResponse, error) {
	data, err := c.c.DoForm("POST", "/audio/translations", req)
	if err != nil {
		return nil, err
	}
	var resp TranslateResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
