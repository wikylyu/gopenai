# GOpenAI

Go implemention for [OpenAI API](https://platform.openai.com/docs/api-reference/introduction)

GOpenAPI provides their official python-like api. 


## Usage

### create completion

```golang
package main

import (
	"fmt"
	"os"

	"github.com/wikylyu/gopenai"
	"github.com/wikylyu/gopenai/completions"
)

func main() {
	openai := gopenai.New(&gopenai.Config{
		ApiKey: os.Getenv("OPENAI_API_KEY"),
	})

	prompt := "Say this is a test"
	resp, err := openai.Completion.Create(&completions.CreateRequest{
		Model:       "text-davinci-003",
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
	if apierr:=err.(*client.Error);apierr!=nil{
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
* [ ] Moderations
* [ ] Engines

## Development

This project is in development.