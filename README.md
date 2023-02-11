# GOpenAI
[![GoDoc](http://img.shields.io/badge/GoDoc-Reference-blue.svg)](https://pkg.go.dev/github.com/wikylyu/gopenai)
[![Go Report Card](https://goreportcard.com/badge/github.com/wikylyu/gopenai)](https://goreportcard.com/report/github.com/wikylyu/gopenai)

Unofficial Go implemention for [OpenAI API](https://platform.openai.com/docs/api-reference/introduction), adhering to their official [OpenAPI spec](https://github.com/openai/openai-openapi/blob/master/openapi.yaml).

GOpenAI provides their python-style like api. It's easy to understand and use.

* openai.Completion.Create(...)
* openai.Edit.Create(...)
* ...

## Usage

### Create Completion

```golang
package main

import (
	"fmt"
	"os"

	"github.com/wikylyu/gopenai"
	"github.com/wikylyu/gopenai/completions"
	"github.com/wikylyu/gopenai/models"
)

func main() {
	openai := gopenai.New(&gopenai.Config{
		ApiKey: os.Getenv("OPENAI_API_KEY"),
	})

	prompt := []string{"Say this is a test"}
	resp, err := openai.Completion.Create(&completions.CreateRequest{
		Model:       models.ModelTextDavinci003,
		Prompt:      prompt,
		MaxTokens:   256,
		Temperature: 0,
	})
	if err != nil {
		panic(err)
	}
	for _, choice := range resp.Choices {
		fmt.Printf("%s\n", choice.Text)
	}
}

```

### For more usages, see ***examples/*** folder.

## Error Handling

There're two kinds of errors that api method may return.

1. OpenAI api error
2. Network or general error

You can use following code to process errors.

```golang

resp,err:=openai.Completion.Create(...)
if err!=nil{
	if apierr:=err.(*api.Error);apierr!=nil{
		/* OpenAI api error, read apierr.Message or apierr.Type to determine exact error reason */
	}else {
		/* Network error */
	}
}

```


## API List

* [x] Model
* [x] Completion
* [x] Edit
* [x] Images
* [x] Embeddings
* [x] Files
* [x] Fine-tunes
* [x] Moderations
* [ ] Engines

## Development

[Engines](https://platform.openai.com/docs/api-reference/engines) is not going to be implemented, cause it's deprecated. [Models](https://platform.openai.com/docs/api-reference/models) is their replacement.
