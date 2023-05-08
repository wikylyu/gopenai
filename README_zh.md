# GOpenAI
[![GoDoc](http://img.shields.io/badge/GoDoc-Reference-blue.svg)](https://pkg.go.dev/github.com/wikylyu/gopenai)
[![Go Report Card](https://goreportcard.com/badge/github.com/wikylyu/gopenai)](https://goreportcard.com/report/github.com/wikylyu/gopenai)

GOpenAI 是一个非官方的[OpenAI API](https://platform.openai.com/docs/api-reference/introduction)实现, 遵循官方的[OpenAPI 规范](https://github.com/openai/openai-openapi/blob/master/openapi.yaml).

GOpenAI提供了类似于Python风格的API，易于理解和使用。

* openai.Completion.Create(...)
* openai.Edit.Create(...)
* ...

## 用法

### 创建 Completion

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

### 更多用法, 请参见 ***examples/*** 目录.

## 错误处理

API方法可能返回两种类型的错误。

1. OpenAI API错误
2. 网络或通用错误

您可以使用以下代码来处理错误。

```golang

resp,err:=openai.Completion.Create(...)
if err!=nil{
	if apierr:=err.(*api.Error);apierr!=nil{
		/* OpenAI api错误，请读取apierr.Message或apierr.Type以确定确切的错误原因 */
	}else {
		/* 网络错误 */
	}
}

```

### 命令行工具


更多详情请参见 **[tool/](https://github.com/wikylyu/gopenai/tree/main/tool)** 目录。


## API List

* [x] Model
  * [x] Create
  * [x] Retrieve
  * [x] Delete
* [x] Completion
  * [x] Create
* [x] Chat
  * [x] Create
* [x] Edit
  * [x] Create
* [x] Images
  * [x] Create
  * [x] Edit
  * [x] Variation
* [x] Embeddings
  * [x] Create
* [x] Audio
  * [x] Transcribe
  * [x] Translate
* [x] Files
  * [x] Create
  * [x] Retrieve
  * [x] Download
* [x] Fine-tunes
  * [x] Create
  * [x] Retrieve
  * [x] List
  * [x] Cancel
  * [x] Events
* [x] Moderations
  * [x] Create

* [ ] Engines

## Development

[Engines](https://platform.openai.com/docs/api-reference/engines) 已经被官方废弃，不会实现； 请使用[Models](https://platform.openai.com/docs/api-reference/models)作为替代。
