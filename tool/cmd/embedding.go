package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/wikylyu/gopenai/embeddings"
)

var EmbeddingCmd = &cobra.Command{
	Use:   "embedding",
	Short: "Get a vector representation of a given input that can be easily consumed by machine learning models and algorithms.",
	Long:  "Get a vector representation of a given input that can be easily consumed by machine learning models and algorithms.",

	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var embeddingCreateParam struct {
	Model string
	Input []string
	User  string
}

var embeddingCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates an embedding vector representing the input text.",
	Long:  "Creates an embedding vector representing the input text.",
	Run: func(cmd *cobra.Command, args []string) {
		model := embeddingCreateParam.Model
		input := embeddingCreateParam.Input
		user := embeddingCreateParam.User

		r, err := openai.Embedding.Create(&embeddings.CreateRequest{
			Model: model,
			Input: input,
			User:  user,
		})
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			return
		}
		printJson(r)
	},
}

func init() {

	embeddingCreateCmd.PersistentFlags().StringVarP(&embeddingCreateParam.Model, "model", "m", "", "ID of the model to use. You can use the List models API to see all of your available models, or see our Model overview for descriptions of them.")
	embeddingCreateCmd.PersistentFlags().StringArrayVarP(&embeddingCreateParam.Input, "input", "i", nil, "Input text to get embeddings for, encoded as a string or array of tokens. To get embeddings for multiple inputs in a single request, pass an array of strings or array of token arrays. Each input must not exceed 8192 tokens in length.")
	embeddingCreateCmd.PersistentFlags().StringVarP(&embeddingCreateParam.User, "user", "u", "", "A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse.")
	embeddingCreateCmd.MarkPersistentFlagRequired("model")
	embeddingCreateCmd.MarkPersistentFlagRequired("input")

	EmbeddingCmd.AddCommand(embeddingCreateCmd)
}
