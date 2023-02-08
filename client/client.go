package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	endpoint string
	apikey   string

	httpClient *http.Client
}

const DefaultEndpoint = "https://api.openai.com/v1"

func NewClient(endpoint, apikey string) *Client {
	tr := &http.Transport{
		MaxIdleConns:    10,
		IdleConnTimeout: 30 * time.Second,
	}
	c := &Client{
		endpoint:   endpoint,
		apikey:     apikey,
		httpClient: &http.Client{Transport: tr},
	}
	if c.endpoint == "" {
		c.endpoint = DefaultEndpoint
	}

	return c
}

func (c *Client) getRequest(method, path string, body interface{}) (*http.Request, error) {
	fullurl, err := url.JoinPath(c.endpoint, path)
	if err != nil {
		return nil, err
	}
	var reader io.Reader = nil
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		reader = bytes.NewBuffer(data)
	}
	req, err := http.NewRequest(method, fullurl, reader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apikey))
	return req, nil
}

func (c *Client) Do(method, path string, body interface{}) ([]byte, error) {
	req, err := c.getRequest(method, path, body)
	if err != nil {
		return nil, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		var e ErrorResponse
		if err := json.Unmarshal(data, &e); err != nil {
			return nil, err
		}
		return nil, &e.Error
	}

	return data, nil
}
