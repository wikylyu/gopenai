package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ModelCmd = &cobra.Command{
	Use:  "model",
	Long: "List and describe the various models available in the API. You can refer to the Models documentation to understand what models are available and the differences between them.",

	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var listModelCmd = &cobra.Command{
	Use:  "list",
	Long: "Lists the currently available models, and provides basic information about each one such as the owner and availability.",
	Run: func(cmd *cobra.Command, args []string) {
		r, err := openai.Model.List()
		if err != nil {
			fmt.Printf("error: %v\n", err)
			return
		}
		printJson(r)
	},
}

var retriveModelParams struct {
	Model string
}

var retriveModelCmd = &cobra.Command{
	Use:  "retrieve",
	Long: "Retrieves a model instance, providing basic information about the model such as the owner and permissioning.",
	Run: func(cmd *cobra.Command, args []string) {
		model := retriveModelParams.Model
		r, err := openai.Model.Retrieve(model)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			return
		}
		printJson(r)
	},
}

var deleteModelParams struct {
	Model string
}

var deleteModelCmd = &cobra.Command{
	Use:  "delete",
	Long: "Delete a fine-tuned model. You must have the Owner role in your organization.",
	Run: func(cmd *cobra.Command, args []string) {
		model := deleteModelParams.Model
		r, err := openai.Model.Delete(model)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			return
		}
		printJson(r)
	},
}

func init() {
	retriveModelCmd.PersistentFlags().StringVarP(&retriveModelParams.Model, "model", "m", "", "[required] The ID of the model to use for this request")
	retriveModelCmd.MarkPersistentFlagRequired("model")

	deleteModelCmd.PersistentFlags().StringVarP(&deleteModelParams.Model, "model", "m", "", "[required] The model to delete")
	deleteModelCmd.MarkPersistentFlagRequired("model")

	ModelCmd.AddCommand(listModelCmd, retriveModelCmd, deleteModelCmd)

}
