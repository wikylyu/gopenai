package main

import (
	"github.com/spf13/cobra"
	"github.com/wikylyu/gopenai/tool/cmd"
)

var rootCmd = &cobra.Command{
	Use:   "gopenai",
	Short: "A command line tool for OpenAI",
	Long:  "A command line tool for OpenAI.\nThis tool will load OpenAI api key from environment variable.\n\n\tOPENAI_API_KEY=yourkey gopenai ...",

	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(cmd.CompletionCmd, cmd.ModelCmd)
}

func main() {
	rootCmd.Execute()

}
