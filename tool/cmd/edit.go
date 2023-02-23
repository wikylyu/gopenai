package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/wikylyu/gopenai/edits"
)

var EditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Given a prompt and an instruction, the model will return an edited version of the prompt.",
	Long:  "Given a prompt and an instruction, the model will return an edited version of the prompt.",

	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var createEditParams struct {
	Model       string
	Input       string
	Instruction string
	N           int64
	Temperature float64
	TopP        float64
}

var createEditCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a new edit for the provided input, instruction, and parameters.",
	Long:  "Creates a new edit for the provided input, instruction, and parameters.",
	Run: func(cmd *cobra.Command, args []string) {
		model := createEditParams.Model
		input := createEditParams.Input
		instruction := createEditParams.Instruction
		n := createEditParams.N
		temperature := createEditParams.Temperature
		topP := createEditParams.TopP
		r, err := openai.Edit.Create(&edits.CreateRequest{
			Model:       model,
			Input:       input,
			Instruction: instruction,
			N:           n,
			Temperature: temperature,
			TopP:        topP,
		})
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			return
		}
		printJson(r)
	},
}

func init() {
	createEditCmd.PersistentFlags().StringVarP(&createEditParams.Model, "model", "m", "", "ID of the model to use. You can use the text-davinci-edit-001 or code-davinci-edit-001 model with this endpoint.")
	createEditCmd.PersistentFlags().StringVarP(&createEditParams.Input, "input", "p", "", "The input text to use as a starting point for the edit.")
	createEditCmd.PersistentFlags().StringVarP(&createEditParams.Instruction, "instruction", "i", "", "The instruction that tells the model how to edit the prompt.")
	createEditCmd.PersistentFlags().Int64VarP(&createEditParams.N, "n", "n", 1, "How many edits to generate for the input and instruction.")
	createEditCmd.PersistentFlags().Float64VarP(&createEditParams.Temperature, "temperature", "t", 1, "What sampling temperature to use, between 0 and 2. Higher values like 0.8 will make the output more random, while lower values like 0.2 will make it more focused and deterministic.\nWe generally recommend altering this or top_p but not both.")
	createEditCmd.PersistentFlags().Float64VarP(&createEditParams.TopP, "top-p", "", 1, "An alternative to sampling with temperature, called nucleus sampling, where the model considers the results of the tokens with top_p probability mass. So 0.1 means only the tokens comprising the top 10% probability mass are considered.\nWe generally recommend altering this or temperature but not both.")
	createEditCmd.MarkPersistentFlagRequired("model")
	createEditCmd.MarkPersistentFlagRequired("instruction")
	EditCmd.AddCommand(createEditCmd)
}
