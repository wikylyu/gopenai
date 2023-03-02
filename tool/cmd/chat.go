package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/wikylyu/gopenai/chat"
)

var ChatCommand = &cobra.Command{
	Use:   "chat",
	Short: "Given a chat conversation, the model will return a chat completion response.",
	Long:  "Given a chat conversation, the model will return a chat completion response.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var chatCreateParams struct {
	Model            string
	Messages         string
	Temperature      float64
	TopP             float64
	N                int64
	MaxTokens        int64
	Stream           bool
	Stop             []string
	PresencePenalty  float64
	FrequencyPenalty float64
	LogitBias        string
	User             string
}

var chatCreateCommand = &cobra.Command{
	Use:   "create",
	Short: "Creates a completion for the chat message",
	Long:  "Creates a completion for the chat message",
	Run: func(cmd *cobra.Command, args []string) {
		model := chatCreateParams.Model
		messages := chatCreateParams.Messages
		temperature := chatCreateParams.Temperature
		topP := chatCreateParams.TopP
		maxTokens := chatCreateParams.MaxTokens
		n := chatCreateParams.N
		stream := chatCreateParams.Stream
		stop := chatCreateParams.Stop
		presencePenalty := chatCreateParams.PresencePenalty
		frequencyPenalty := chatCreateParams.FrequencyPenalty
		logitBias := chatCreateParams.LogitBias
		user := chatCreateParams.User
		logitBiasMap := make(map[string]interface{})
		if logitBias != "" {
			if err := json.Unmarshal([]byte(logitBias), &logitBiasMap); err != nil {
				fmt.Printf("error: failed to parse logit-bias: %v\n", err)
				return
			}
		}
		messagesData := make([]*chat.Message, 0)
		if err := json.Unmarshal([]byte(messages), &messagesData); err != nil {
			fmt.Printf("error: failed to parse message: %v\n", err)
			return
		}

		r, err := openai.Chat.Create(&chat.CreateRequest{
			Model:            model,
			Messages:         messagesData,
			Temperature:      temperature,
			TopP:             topP,
			N:                n,
			Stream:           stream,
			Stop:             stop,
			PresencePenalty:  presencePenalty,
			FrequencyPenalty: frequencyPenalty,
			User:             user,
			LogitBias:        logitBiasMap,
			MaxTokens:        maxTokens,
		})
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			return
		}
		printJson(r)
	},
}

func init() {
	chatCreateCommand.PersistentFlags().StringVarP(&chatCreateParams.Model, "model", "m", "", "ID of the model to use. Currently, only gpt-3.5-turbo and gpt-3.5-turbo-0301 are supported.")
	chatCreateCommand.PersistentFlags().StringVarP(&chatCreateParams.Messages, "messages", "", "", "The messages to generate chat completions for, in the chat format.")
	chatCreateCommand.PersistentFlags().Int64VarP(&chatCreateParams.MaxTokens, "max-tokens", "", 0, "The maximum number of tokens allowed for the generated answer. By default, the number of tokens the model can return will be (4096 - prompt tokens).")
	chatCreateCommand.PersistentFlags().Float64VarP(&chatCreateParams.Temperature, "temperature", "", 1, "What sampling temperature to use, between 0 and 2. Higher values like 0.8 will make the output more random, while lower values like 0.2 will make it more focused and deterministic.\nWe generally recommend altering this or top_p but not both.")
	chatCreateCommand.PersistentFlags().Float64VarP(&chatCreateParams.TopP, "top-p", "", 1, "An alternative to sampling with temperature, called nucleus sampling, where the model considers the results of the tokens with top_p probability mass. So 0.1 means only the tokens comprising the top 10% probability mass are considered.\nWe generally recommend altering this or temperature but not both.")
	chatCreateCommand.PersistentFlags().Int64VarP(&chatCreateParams.N, "n", "n", 1, "How many completions to generate for each prompt.\nNote: Because this parameter generates many completions, it can quickly consume your token quota. Use carefully and ensure that you have reasonable settings for max_tokens and stop.")
	chatCreateCommand.PersistentFlags().BoolVarP(&chatCreateParams.Stream, "stream", "", false, "Whether to stream back partial progress. If set, tokens will be sent as data-only server-sent events as they become available, with the stream terminated by a data: [DONE] message.")
	chatCreateCommand.PersistentFlags().StringArrayVarP(&chatCreateParams.Stop, "stop", "", nil, "Up to 4 sequences where the API will stop generating further tokens. The returned text will not contain the stop sequence.")
	chatCreateCommand.PersistentFlags().Float64VarP(&chatCreateParams.PresencePenalty, "presence-penalty", "", 0, "Number between -2.0 and 2.0. Positive values penalize new tokens based on whether they appear in the text so far, increasing the model's likelihood to talk about new topics.")
	chatCreateCommand.PersistentFlags().Float64VarP(&chatCreateParams.FrequencyPenalty, "frequency-penalty", "", 0, "Number between -2.0 and 2.0. Positive values penalize new tokens based on their existing frequency in the text so far, decreasing the model's likelihood to repeat the same line verbatim.")
	chatCreateCommand.PersistentFlags().StringVarP(&chatCreateParams.LogitBias, "logit-bias", "", "", "Modify the likelihood of specified tokens appearing in the completion.\nAccepts a json object that maps tokens (specified by their token ID in the GPT tokenizer) to an associated bias value from -100 to 100. You can use this tokenizer tool (which works for both GPT-2 and GPT-3) to convert text to token IDs. Mathematically, the bias is added to the logits generated by the model prior to sampling. The exact effect will vary per model, but values between -1 and 1 should decrease or increase likelihood of selection; values like -100 or 100 should result in a ban or exclusive selection of the relevant token.\nAs an example, you can pass {\"50256\": -100} to prevent the <|endoftext|> token from being generated.")
	chatCreateCommand.PersistentFlags().StringVarP(&chatCreateParams.User, "user", "u", "", "A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse.")
	chatCreateCommand.MarkPersistentFlagRequired("model")
	chatCreateCommand.MarkPersistentFlagRequired("messages")

	ChatCommand.AddCommand(chatCreateCommand)
}
