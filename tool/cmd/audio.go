package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/wikylyu/gopenai/audio"
)

var AudioCmd = &cobra.Command{
	Use:   "audio",
	Short: "Learn how to turn audio into text.",
	Long:  "Learn how to turn audio into text.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var audioTranscribeParams struct {
	File           string
	Model          string
	Prompt         string
	ResponseFormat string
	Temperature    float64
	Language       string
}

var audioTranscribeCmd = &cobra.Command{
	Use:   "transcribe",
	Short: "Transcribes audio into the input language.",
	Long:  "Transcribes audio into the input language.",
	Run: func(cmd *cobra.Command, args []string) {
		file := audioTranscribeParams.File
		model := audioTranscribeParams.Model
		prompt := audioTranscribeParams.Prompt
		responseFormat := audioTranscribeParams.ResponseFormat
		tempeature := audioTranscribeParams.Temperature
		language := audioTranscribeParams.Language

		f, err := os.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			return
		}
		defer f.Close()

		r, err := openai.Audio.Transcribe(&audio.TranscribeRequest{
			File:           f,
			Model:          model,
			Prompt:         prompt,
			ResponseFormat: responseFormat,
			Temperature:    tempeature,
			Language:       language,
		})
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			return
		}
		printJson(r)
	},
}

var audioTranslateParams struct {
	File           string
	Model          string
	Prompt         string
	ResponseFormat string
	Temperature    float64
}

var audioTranslateCmd = &cobra.Command{
	Use:   "translate",
	Short: "Translates audio into into English.",
	Long:  "Translates audio into into English.",
	Run: func(cmd *cobra.Command, args []string) {
		file := audioTranslateParams.File
		model := audioTranslateParams.Model
		prompt := audioTranslateParams.Prompt
		responseFormat := audioTranslateParams.ResponseFormat
		tempeature := audioTranslateParams.Temperature

		f, err := os.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			return
		}
		defer f.Close()

		r, err := openai.Audio.Translate(&audio.TranslateRequest{
			File:           f,
			Model:          model,
			Prompt:         prompt,
			ResponseFormat: responseFormat,
			Temperature:    tempeature,
		})
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			return
		}
		printJson(r)
	},
}

func init() {
	audioTranscribeCmd.PersistentFlags().StringVarP(&audioTranscribeParams.File, "file", "f", "", "The audio file to transcribe, in one of these formats: mp3, mp4, mpeg, mpga, m4a, wav, or webm.")
	audioTranscribeCmd.PersistentFlags().StringVarP(&audioTranscribeParams.Model, "model", "m", "", "ID of the model to use. Only whisper-1 is currently available.")
	audioTranscribeCmd.PersistentFlags().StringVarP(&audioTranscribeParams.Prompt, "prompt", "p", "", "An optional text to guide the model's style or continue a previous audio segment. The prompt should match the audio language.")
	audioTranscribeCmd.PersistentFlags().StringVarP(&audioTranscribeParams.ResponseFormat, "response-format", "r", "json", "The format of the transcript output, in one of these options: json, text, srt, verbose_json, or vtt.")
	audioTranscribeCmd.PersistentFlags().Float64VarP(&audioTranscribeParams.Temperature, "temperature", "t", 0, "The sampling temperature, between 0 and 1. Higher values like 0.8 will make the output more random, while lower values like 0.2 will make it more focused and deterministic. If set to 0, the model will use log probability to automatically increase the temperature until certain thresholds are hit.")
	audioTranscribeCmd.PersistentFlags().StringVarP(&audioTranscribeParams.Language, "language", "l", "", "The language of the input audio. Supplying the input language in ISO-639-1 format will improve accuracy and latency.")
	audioTranscribeCmd.MarkPersistentFlagRequired("file")
	audioTranscribeCmd.MarkPersistentFlagRequired("model")

	audioTranslateCmd.PersistentFlags().StringVarP(&audioTranslateParams.File, "file", "f", "", "The audio file to transcribe, in one of these formats: mp3, mp4, mpeg, mpga, m4a, wav, or webm.")
	audioTranslateCmd.PersistentFlags().StringVarP(&audioTranslateParams.Model, "model", "m", "", "ID of the model to use. Only whisper-1 is currently available.")
	audioTranslateCmd.PersistentFlags().StringVarP(&audioTranslateParams.Prompt, "prompt", "p", "", "An optional text to guide the model's style or continue a previous audio segment. The prompt should match the audio language.")
	audioTranslateCmd.PersistentFlags().StringVarP(&audioTranslateParams.ResponseFormat, "response-format", "r", "json", "The format of the transcript output, in one of these options: json, text, srt, verbose_json, or vtt.")
	audioTranslateCmd.PersistentFlags().Float64VarP(&audioTranslateParams.Temperature, "temperature", "t", 0, "The sampling temperature, between 0 and 1. Higher values like 0.8 will make the output more random, while lower values like 0.2 will make it more focused and deterministic. If set to 0, the model will use log probability to automatically increase the temperature until certain thresholds are hit.")
	audioTranslateCmd.MarkPersistentFlagRequired("file")
	audioTranslateCmd.MarkPersistentFlagRequired("model")

	AudioCmd.AddCommand(audioTranscribeCmd, audioTranslateCmd)
}
