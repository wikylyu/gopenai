package main

import (
	"fmt"
	"io"
	"os"

	"github.com/wikylyu/gopenai"
	"github.com/wikylyu/gopenai/files"
)

func main() {
	openai := gopenai.New(&gopenai.Config{
		ApiKey: os.Getenv("OPENAI_API_KEY"),
	})

	file, err := os.Open("./mydata.jsonl")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	resp, err := openai.Files.Create(&files.CreateRequest{
		File:    file,
		Purpose: "fine-tune",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("uploaded: %s:%s\n", resp.ID, resp.Filename)
	fmt.Printf("--------------------------------------------\n")

	resp2, err := openai.Files.List()
	if err != nil {
		panic(err)
	}
	for _, file := range resp2.Data {
		fmt.Printf("%s:%s\n", file.ID, file.Filename)
	}
	fmt.Printf("--------------------------------------------\n")

	resp3, err := openai.Files.Retrieve(resp.ID)
	if err != nil {
		panic(err)
	}
	if resp3.Filename != resp.Filename || resp3.Bytes != resp.Bytes {
		panic("BUG")
	}

	fmt.Printf("--------------------------------------------\n")
	reader, err := openai.Files.Download(resp.ID)
	if err != nil {
		panic(err)
	}
	defer reader.Close()
	content, err := io.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", string(content))
}
