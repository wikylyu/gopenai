package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wikylyu/gopenai/completions"
)

var CompletionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Text completion",
	Long:  "Given a prompt, the model will return one or more predicted completions, and can also return the probabilities of alternative tokens at each position.",

	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var completionCreateParams struct {
	Model            string
	Prompt           []string
	Suffix           string
	MaxTokens        int64
	Temperature      float64
	TopP             float64
	N                int64
	Stream           bool
	Logprobs         int64
	Echo             bool
	Stop             []string
	PresencePenalty  float64
	FrequencyPenalty float64
	BestOf           int64
	LogitBias        string
	User             string
}

var completionCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a completion",
	Long:  "Creates a completion for the provided prompt and parameters",

	Run: func(cmd *cobra.Command, args []string) {
		params := completionCreateParams
		if params.Model == "" {
			fmt.Printf("model is required\n")
			return
		}
		logitBias := make(map[string]interface{})
		if params.LogitBias != "" {
			if err := json.Unmarshal([]byte(params.LogitBias), &logitBias); err != nil {
				fmt.Printf("error: failed to parse logit-bias: %v\n", err)
				return
			}
		}
		r, err := openai.Completion.Create(&completions.CreateRequest{
			Model:            params.Model,
			Prompt:           params.Prompt,
			Suffix:           params.Suffix,
			MaxTokens:        params.MaxTokens,
			Temperature:      params.Temperature,
			TopP:             params.TopP,
			N:                params.N,
			Stream:           params.Stream,
			Logprobs:         params.Logprobs,
			Echo:             params.Echo,
			Stop:             params.Stop,
			PresencePenalty:  params.PresencePenalty,
			FrequencyPenalty: params.FrequencyPenalty,
			BestOf:           params.BestOf,
			LogitBias:        logitBias,
			User:             params.User,
		})
		if err != nil {
			fmt.Printf("error: %v\n", err)
			return
		}
		result, err := json.MarshalIndent(r, "", "  ")
		if err != nil {
			fmt.Printf("error: %v\n", err)
			return
		}
		fmt.Printf("%v\n", string(result))
	},
}

func init() {
	completionCreateCmd.PersistentFlags().StringVarP(&completionCreateParams.Model, "model", "m", "", "ID of the model to use.\nYou can use the List models API to see all of your available models, or see our Model overview for descriptions of them.")
	completionCreateCmd.PersistentFlags().StringArrayVarP(&completionCreateParams.Prompt, "prompt", "p", nil, "The prompt(s) to generate completions for, encoded as a string, array of strings, array of tokens, or array of token arrays.\nNote that <|endoftext|> is the document separator that the model sees during training, so if a prompt is not specified the model will generate as if from the beginning of a new document.")
	completionCreateCmd.PersistentFlags().StringVarP(&completionCreateParams.Suffix, "suffix", "s", "", "The suffix that comes after a completion of inserted text.")
	completionCreateCmd.PersistentFlags().Int64VarP(&completionCreateParams.MaxTokens, "max-tokens", "", 16, "The maximum number of tokens to generate in the completion.\nThe token count of your prompt plus max_tokens cannot exceed the model's context length. Most models have a context length of 2048 tokens (except for the newest models, which support 4096).")
	completionCreateCmd.PersistentFlags().Float64VarP(&completionCreateParams.Temperature, "temperature", "", 1, "What sampling temperature to use, between 0 and 2. Higher values like 0.8 will make the output more random, while lower values like 0.2 will make it more focused and deterministic.\nWe generally recommend altering this or top_p but not both.")
	completionCreateCmd.PersistentFlags().Float64VarP(&completionCreateParams.TopP, "top-p", "", 1, "An alternative to sampling with temperature, called nucleus sampling, where the model considers the results of the tokens with top_p probability mass. So 0.1 means only the tokens comprising the top 10% probability mass are considered.\nWe generally recommend altering this or temperature but not both.")
	completionCreateCmd.PersistentFlags().Int64VarP(&completionCreateParams.N, "n", "n", 1, "How many completions to generate for each prompt.\nNote: Because this parameter generates many completions, it can quickly consume your token quota. Use carefully and ensure that you have reasonable settings for max_tokens and stop.")
	completionCreateCmd.PersistentFlags().BoolVarP(&completionCreateParams.Stream, "stream", "", false, "Whether to stream back partial progress. If set, tokens will be sent as data-only server-sent events as they become available, with the stream terminated by a data: [DONE] message.")
	completionCreateCmd.PersistentFlags().Int64VarP(&completionCreateParams.Logprobs, "logprobs", "l", 0, "Include the log probabilities on the logprobs most likely tokens, as well the chosen tokens. For example, if logprobs is 5, the API will return a list of the 5 most likely tokens. The API will always return the logprob of the sampled token, so there may be up to logprobs+1 elements in the response.\nThe maximum value for logprobs is 5. If you need more than this, please contact us through our Help center and describe your use case.")
	completionCreateCmd.PersistentFlags().BoolVarP(&completionCreateParams.Echo, "echo", "e", false, "Echo back the prompt in addition to the completion")
	completionCreateCmd.PersistentFlags().StringArrayVarP(&completionCreateParams.Stop, "stop", "", nil, "Up to 4 sequences where the API will stop generating further tokens. The returned text will not contain the stop sequence.")
	completionCreateCmd.PersistentFlags().Float64VarP(&completionCreateParams.PresencePenalty, "presence-penalty", "", 0, "Number between -2.0 and 2.0. Positive values penalize new tokens based on whether they appear in the text so far, increasing the model's likelihood to talk about new topics.")
	completionCreateCmd.PersistentFlags().Float64VarP(&completionCreateParams.FrequencyPenalty, "frequency-penalty", "", 0, "Number between -2.0 and 2.0. Positive values penalize new tokens based on their existing frequency in the text so far, decreasing the model's likelihood to repeat the same line verbatim.")
	completionCreateCmd.PersistentFlags().Int64VarP(&completionCreateParams.BestOf, "best-of", "", 1, "Generates best_of completions server-side and returns the \"best\" (the one with the highest log probability per token). Results cannot be streamed.\nWhen used with n, best_of controls the number of candidate completions and n specifies how many to return – best_of must be greater than n.\nNote: Because this parameter generates many completions, it can quickly consume your token quota. Use carefully and ensure that you have reasonable settings for max_tokens and stop.")
	completionCreateCmd.PersistentFlags().StringVarP(&completionCreateParams.LogitBias, "logit-bias", "", "", "Modify the likelihood of specified tokens appearing in the completion.\nAccepts a json object that maps tokens (specified by their token ID in the GPT tokenizer) to an associated bias value from -100 to 100. You can use this tokenizer tool (which works for both GPT-2 and GPT-3) to convert text to token IDs. Mathematically, the bias is added to the logits generated by the model prior to sampling. The exact effect will vary per model, but values between -1 and 1 should decrease or increase likelihood of selection; values like -100 or 100 should result in a ban or exclusive selection of the relevant token.\nAs an example, you can pass {\"50256\": -100} to prevent the <|endoftext|> token from being generated.")
	completionCreateCmd.PersistentFlags().StringVarP(&completionCreateParams.User, "user", "u", "", "A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse.")
	CompletionCmd.AddCommand(completionCreateCmd)
}
