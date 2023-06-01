package chat

import (
	"bufio"
	"encoding/json"
	"strings"

	"github.com/wikylyu/gopenai/api"
)

func NewClient(c *api.Client) *ChatClient {
	return &ChatClient{c: c}
}

func (c *ChatClient) Create(req *CreateRequest) (*CreateResponse, error) {
	req.Stream = false
	data, err := c.c.DoJson("POST", "/chat/completions", req)
	if err != nil {
		return nil, err
	}

	var resp CreateResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *ChatClient) CreateStream(req *CreateRequest) (*ChatStreamer, error) {
	req.Stream = true
	resp, err := c.c.DoStream("POST", "/chat/completions", req)
	if err != nil {
		return nil, err
	}

	return &ChatStreamer{
		resp:   resp,
		reader: bufio.NewReader(resp.Body),
		buf:    strings.Builder{},
	}, nil
}
