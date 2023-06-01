package chat

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (s *ChatStreamer) Close() error {
	return s.resp.Body.Close()
}

func (s *ChatStreamer) Read() (*StreamResponse, error) {
read:
	line, err := s.reader.ReadBytes('\n')
	if err != nil {
		return nil, err
	}

	headerData := []byte("data: ")
	line = bytes.TrimSpace(line)
	if len(line) == 0 {
		goto read
	}
	if !bytes.HasPrefix(line, headerData) {
		return nil, fmt.Errorf("invalid data")
	}

	line = bytes.TrimPrefix(line, headerData)
	if string(line) == "[DONE]" {
		return nil, nil
	}
	var resp StreamResponse
	if err := json.Unmarshal(line, &resp); err != nil {
		return nil, err
	}
	for _, choice := range resp.Choices {
		if choice.Delta.Content != "" {
			s.buf.Write([]byte(choice.Delta.Content))
		}
	}
	return &resp, nil
}

/*
 * Get all content data
 */
func (s *ChatStreamer) Content() string {
	return s.buf.String()
}
