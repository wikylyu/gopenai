package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/wikylyu/gopenai/images"
)

var ImageCommand = &cobra.Command{
	Use:  "image",
	Long: "Given a prompt and/or an input image, the model will generate a new image.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var imageCreateParams struct {
	Prompt          string
	N               int64
	Size            string
	RespoinseFormat string
	User            string
}

var imageCreateCommand = &cobra.Command{
	Use:  "create",
	Long: "Creates an image given a prompt.",
	Run: func(cmd *cobra.Command, args []string) {
		prompt := imageCreateParams.Prompt
		n := imageCreateParams.N
		size := imageCreateParams.Size
		responseFormat := imageCreateParams.RespoinseFormat
		user := imageCreateParams.User
		r, err := openai.Image.Create(&images.CreateRequest{
			Prompt:         prompt,
			N:              n,
			Size:           size,
			ResponseFormat: responseFormat,
			User:           user,
		})
		if err != nil {
			fmt.Printf("error: %v\n", err)
			return
		}
		printJson(r)
	},
}

var imageCreateEditParams struct {
	Image          string
	Mask           string
	Prompt         string
	N              int64
	Size           string
	ResponseFormat string
	User           string
}

var imageCreateEditCommand = &cobra.Command{
	Use:  "create-edit",
	Long: "Creates an edited or extended image given an original image and a prompt.",
	Run: func(cmd *cobra.Command, args []string) {
		image := imageCreateEditParams.Image
		mask := imageCreateEditParams.Mask
		prompt := imageCreateEditParams.Prompt
		n := imageCreateEditParams.N
		size := imageCreateEditParams.Size
		responseFormat := imageCreateEditParams.ResponseFormat
		user := imageCreateEditParams.User

		var imageFile, maskFile *os.File
		var err error

		imageFile, err = os.Open(image)
		if err != nil {
			fmt.Printf("open file error: %v\n", err)
			return
		}
		if mask != "" {
			maskFile, err = os.Open(mask)
			fmt.Printf("open file error: %v\n", err)
			return
		}

		r, err := openai.Image.CreateEdit(&images.CreateEditRequest{
			Image:          imageFile,
			Mask:           maskFile,
			Prompt:         prompt,
			N:              n,
			Size:           size,
			ResponseFormat: responseFormat,
			User:           user,
		})
		if err != nil {
			fmt.Printf("error: %v\n", err)
			return
		}
		printJson(r)
	},
}

var imageCreateVariationParams struct {
	Image          string
	N              int64
	Size           string
	ResponseFormat string
	User           string
}

var imageCreateVariationCommand = &cobra.Command{
	Use:  "create-variation",
	Long: "Creates a variation of a given image.",
	Run: func(cmd *cobra.Command, args []string) {
		image := imageCreateVariationParams.Image
		n := imageCreateVariationParams.N
		size := imageCreateVariationParams.Size
		responseFormat := imageCreateVariationParams.ResponseFormat
		user := imageCreateVariationParams.User

		imageFile, err := os.Open(image)
		if err != nil {
			fmt.Printf("open file error: %v\n", err)
			return
		}

		r, err := openai.Image.CreateVariation(&images.CreateVariationRequest{
			Image:          imageFile,
			N:              n,
			Size:           size,
			ResponseFormat: responseFormat,
			User:           user,
		})
		if err != nil {
			fmt.Printf("error: %v\n", err)
			return
		}
		printJson(r)
	},
}

func init() {
	imageCreateCommand.PersistentFlags().StringVarP(&imageCreateParams.Prompt, "prompt", "p", "", "A text description of the desired image(s). The maximum length is 1000 characters.")
	imageCreateCommand.PersistentFlags().Int64VarP(&imageCreateParams.N, "n", "n", 1, "The number of images to generate. Must be between 1 and 10.")
	imageCreateCommand.PersistentFlags().StringVarP(&imageCreateParams.Size, "size", "s", "1024x1024", "The size of the generated images. Must be one of 256x256, 512x512, or 1024x1024.")
	imageCreateCommand.PersistentFlags().StringVarP(&imageCreateParams.RespoinseFormat, "response-format", "r", "url", "The format in which the generated images are returned. Must be one of url or b64_json.")
	imageCreateCommand.PersistentFlags().StringVarP(&imageCreateParams.User, "user", "u", "", "A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse.")
	imageCreateCommand.MarkPersistentFlagRequired("prompt")

	imageCreateEditCommand.PersistentFlags().StringVarP(&imageCreateEditParams.Image, "image", "i", "", "The image to edit. Must be a valid PNG file, less than 4MB, and square. If mask is not provided, image must have transparency, which will be used as the mask.")
	imageCreateEditCommand.PersistentFlags().StringVarP(&imageCreateEditParams.Mask, "mask", "m", "", "An additional image whose fully transparent areas (e.g. where alpha is zero) indicate where image should be edited. Must be a valid PNG file, less than 4MB, and have the same dimensions as image.")
	imageCreateEditCommand.PersistentFlags().StringVarP(&imageCreateEditParams.Prompt, "prompt", "p", "", "A text description of the desired image(s). The maximum length is 1000 characters.")
	imageCreateEditCommand.PersistentFlags().Int64VarP(&imageCreateEditParams.N, "n", "n", 1, "The number of images to generate. Must be between 1 and 10.")
	imageCreateEditCommand.PersistentFlags().StringVarP(&imageCreateEditParams.Size, "size", "s", "1024x1024", "The size of the generated images. Must be one of 256x256, 512x512, or 1024x1024.")
	imageCreateEditCommand.PersistentFlags().StringVarP(&imageCreateEditParams.ResponseFormat, "response-format", "r", "url", "The format in which the generated images are returned. Must be one of url or b64_json.")
	imageCreateEditCommand.PersistentFlags().StringVarP(&imageCreateEditParams.User, "user", "u", "", "A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse.")
	imageCreateEditCommand.MarkPersistentFlagRequired("image")
	imageCreateEditCommand.MarkPersistentFlagRequired("prompt")

	imageCreateVariationCommand.PersistentFlags().StringVarP(&imageCreateVariationParams.Image, "image", "i", "", "The image to use as the basis for the variation(s). Must be a valid PNG file, less than 4MB, and square.")
	imageCreateVariationCommand.PersistentFlags().Int64VarP(&imageCreateVariationParams.N, "n", "n", 1, "The number of images to generate. Must be between 1 and 10.")
	imageCreateVariationCommand.PersistentFlags().StringVarP(&imageCreateVariationParams.Size, "size", "s", "1024x1024", "The size of the generated images. Must be one of 256x256, 512x512, or 1024x1024.")
	imageCreateVariationCommand.PersistentFlags().StringVarP(&imageCreateVariationParams.ResponseFormat, "response-format", "r", "url", "The format in which the generated images are returned. Must be one of url or b64_json.")
	imageCreateVariationCommand.PersistentFlags().StringVarP(&imageCreateVariationParams.User, "user", "u", "", "A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse.")
	imageCreateVariationCommand.MarkPersistentFlagRequired("image")

	ImageCommand.AddCommand(imageCreateCommand, imageCreateEditCommand, imageCreateVariationCommand)
}
