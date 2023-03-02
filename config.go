package gopenai

import "time"

type Config struct {
	BaseURL         string // openai api endpoint, optional
	ApiKey          string // api key, required
	MaxIdleConns    int
	IdleConnTimeout time.Duration
	HttpsProxy      string
}
