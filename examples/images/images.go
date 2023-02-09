package main

import (
	"fmt"
	"os"
	"path"

	"github.com/wikylyu/gopenai"
	"github.com/wikylyu/gopenai/images"
)

func create() {
	openai := gopenai.New(&gopenai.Config{
		ApiKey: os.Getenv("OPENAI_API_KEY"),
	})

	resp, err := openai.Image.Create(&images.CreateRequest{
		Prompt: "A cute baby sea otter",
		N:      2,
		Size:   "1024x1024",
	})
	if err != nil {
		panic(err)
	}
	for _, data := range resp.Data {
		fmt.Printf("%s\n", data.URL)
	}
}

func createEdit() {
	openai := gopenai.New(&gopenai.Config{
		ApiKey: os.Getenv("OPENAI_API_KEY"),
	})
	homedir, _ := os.UserHomeDir()
	file, err := os.Open(path.Join(homedir, "/Pictures/129Magikarp.png"))
	if err != nil {
		panic(err)
	}
	defer file.Close()
	resp2, err := openai.Image.CreateEdit(&images.CreateEditRequest{
		Image:  file,
		Prompt: "put fish on sea beach",
		N:      2,
		Size:   "1024x1024",
	})
	if err != nil {
		panic(err)
	}
	for _, data := range resp2.Data {
		fmt.Printf("%s\n", data.URL)
	}
}

func createVariation() {
	openai := gopenai.New(&gopenai.Config{
		ApiKey: os.Getenv("OPENAI_API_KEY"),
	})
	homedir, _ := os.UserHomeDir()
	file, err := os.Open(path.Join(homedir, "/Pictures/129Magikarp.png"))
	if err != nil {
		panic(err)
	}
	defer file.Close()
	resp, err := openai.Image.CreateVariation(&images.CreateVariationRequest{
		Image: file,
		N:     2,
		Size:  "1024x1024",
	})
	if err != nil {
		panic(err)
	}
	for _, data := range resp.Data {
		fmt.Printf("%s\n", data.URL)
	}
}

func main() {
	create()
	fmt.Printf("--------------------------------------\n")
	createEdit()
	fmt.Printf("--------------------------------------\n")
	createVariation()
}
