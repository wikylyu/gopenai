package main

import (
	"fmt"
	"os"

	"github.com/wikylyu/gopenai"
	"github.com/wikylyu/gopenai/api"
	"github.com/wikylyu/gopenai/files"
	"github.com/wikylyu/gopenai/finetunes"
)

func create() {
	openai := gopenai.New(&gopenai.Config{
		ApiKey: os.Getenv("OPENAI_API_KEY"),
	})

	file, err := os.Open("./mydata.jsonl")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	resp, err := openai.File.Create(&files.CreateRequest{
		File:    file,
		Purpose: "fine-tune",
	})
	if err != nil {
		panic(err)
	}

	finetune, err := openai.FineTune.Create(&finetunes.CreateRequest{
		TrainingFile: resp.ID,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("id:%s\n", finetune.ID)
	fmt.Printf("model:%s\n", finetune.Model)
	fmt.Printf("result_files:%v\n", finetune.ResultFiles)
}

func list() {
	openai := gopenai.New(&gopenai.Config{
		ApiKey: os.Getenv("OPENAI_API_KEY"),
	})

	resp, err := openai.FineTune.List()
	if err != nil {
		panic(err)
	}
	for _, ft := range resp.Data {
		finetune, err := openai.FineTune.Retrieve(ft.ID)
		if err != nil {
			panic(err)
		}
		fmt.Printf("id:%s\n", finetune.ID)
		fmt.Printf("model:%s\n", finetune.Model)
		fmt.Printf("status:%s\n", finetune.Status)
		fmt.Printf("training_files:%v\n", finetune.TrainingFiles)
		fmt.Printf("result_files:%v\n", finetune.ResultFiles)
		fmt.Printf("fine-tuned model: %v\n", finetune.FineTunedModel)
		events, err := openai.FineTune.ListEvents(finetune.ID)
		if err != nil {
			panic(err)
		}
		for _, event := range events.Data {
			fmt.Printf("\tevent:%s\n", event.Message)
		}
		fmt.Printf("--------------------------------------\n")
	}
}

func cancel() {
	openai := gopenai.New(&gopenai.Config{
		ApiKey: os.Getenv("OPENAI_API_KEY"),
	})

	resp, err := openai.FineTune.List()
	if err != nil {
		panic(err)
	}
	for _, ft := range resp.Data {
		finetune, err := openai.FineTune.Cancel(ft.ID)
		if err != nil {
			if apierr := err.(*api.Error); apierr != nil {
				fmt.Printf("%s:%s\n", apierr.Type, apierr.Code)
				continue
			} else {
				panic(err)
			}
		}
		fmt.Printf("id:%s\n", finetune.ID)
		fmt.Printf("model:%s\n", finetune.Model)
		fmt.Printf("status:%s\n", finetune.Status)
		fmt.Printf("training_files:%v\n", finetune.TrainingFiles)
		fmt.Printf("result_files:%v\n", finetune.ResultFiles)

		fmt.Printf("--------------------------------------\n")

	}
}

func main() {
	// create()
	list()
	// cancel()
}
