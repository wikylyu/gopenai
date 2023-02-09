package main

import (
	"fmt"
	"os"

	"github.com/wikylyu/gopenai"
)

func main() {
	openai := gopenai.New(&gopenai.Config{
		ApiKey: os.Getenv("OPENAI_API_KEY"),
	})
	resp, err := openai.Model.List()
	if err != nil {
		panic(err)
	}
	for _, model := range resp.Data {
		fmt.Printf("-------------model-------------\n")
		fmt.Printf("id:\t\t%s\n", model.ID)
		fmt.Printf("object:\t\t%s\n", model.Object)
		fmt.Printf("owned_by:\t%s\n", model.OwnedBy)
		fmt.Printf("created:\t%d\n", model.Created)
		for _, permission := range model.Permission {
			fmt.Printf("permission id:\t%s\n", permission.ID)
		}
	}

	model, err := openai.Model.Retrieve(resp.Data[0].ID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\n")
	fmt.Printf("-------------retrieved model-------------\n")
	fmt.Printf("id:\t\t%s\n", model.ID)
	fmt.Printf("object:\t\t%s\n", model.Object)
	fmt.Printf("owned_by:\t%s\n", model.OwnedBy)
	fmt.Printf("created:\t%d\n", model.Created)
	for _, permission := range model.Permission {
		fmt.Printf("permission id:\t%s\n", permission.ID)
	}

	// deletedModel, err := openai.Model.Delete("fine tuned model")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("model %s deleted\n", deletedModel.ID)

}
