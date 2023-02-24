package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/wikylyu/gopenai/moderations"
)

var ModerationCmd = &cobra.Command{
	Use:   "moderation",
	Short: "Given a input text, outputs if the model classifies it as violating OpenAI's content policy.",
	Long:  "Given a input text, outputs if the model classifies it as violating OpenAI's content policy.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var moderationCreateParams struct {
	Input []string
	Model string
}

var moderationCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Classifies if text violates OpenAI's Content Policy",
	Long:  "Classifies if text violates OpenAI's Content Policy",
	Run: func(cmd *cobra.Command, args []string) {
		input := moderationCreateParams.Input
		model := moderationCreateParams.Model

		r, err := openai.Moderation.Create(&moderations.CreateRequest{
			Input: input,
			Model: model,
		})
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			return
		}
		printJson(r)
	},
}

func init() {

	moderationCreateCmd.PersistentFlags().StringArrayVarP(&moderationCreateParams.Input, "input", "i", nil, "The input text to classify.")
	moderationCreateCmd.PersistentFlags().StringVarP(&moderationCreateParams.Model, "model", "m", "text-moderation-latest", "Two content moderations models are available: text-moderation-stable and text-moderation-latest.\nThe default is text-moderation-latest which will be automatically upgraded over time. This ensures you are always using our most accurate model. If you use text-moderation-stable, we will provide advanced notice before updating the model. Accuracy of text-moderation-stable may be slightly lower than for text-moderation-latest.")

	moderationCreateCmd.MarkPersistentFlagRequired("input")

	ModerationCmd.AddCommand(moderationCreateCmd)
}
