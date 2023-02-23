package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/wikylyu/gopenai/files"
)

var FileCmd = &cobra.Command{
	Use:   "file",
	Short: "Files are used to upload documents that can be used with features like Fine-tuning.",
	Long:  "Files are used to upload documents that can be used with features like Fine-tuning.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var fileCreateParams struct {
	File    string
	Purpose string
}

var fileCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Upload a file.",
	Long:  "Upload a file that contains document(s) to be used across various endpoints/features. Currently, the size of all the files uploaded by one organization can be up to 1 GB. Please contact us if you need to increase the storage limit.",
	Run: func(cmd *cobra.Command, args []string) {
		file := fileCreateParams.File
		purpose := fileCreateParams.Purpose

		f, err := os.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			return
		}
		r, err := openai.File.Create(&files.CreateRequest{
			File:    f,
			Purpose: purpose,
		})
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			return
		}
		printJson(r)
	},
}

var fileDeleteParams struct {
	FileID string
}

var fileDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a file.",
	Long:  "Delete a file.",
	Run: func(cmd *cobra.Command, args []string) {
		fileid := fileDeleteParams.FileID

		r, err := openai.File.Delete(fileid)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			return
		}
		printJson(r)
	},
}

var fileRetrieveParams struct {
	FileID string
}

var fileRetrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Returns information about a specific file.",
	Long:  "Returns information about a specific file.",
	Run: func(cmd *cobra.Command, args []string) {
		fileid := fileRetrieveParams.FileID

		r, err := openai.File.Retrieve(fileid)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			return
		}
		printJson(r)
	},
}

var fileDownloadParams struct {
	FileID string
}

var fileDownloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Returns the contents of the specified file.",
	Long:  "Returns the contents of the specified file.",
	Run: func(cmd *cobra.Command, args []string) {
		fileid := fileDownloadParams.FileID

		r, err := openai.File.Download(fileid)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			return
		}
		defer r.Close()

		buf := make([]byte, 4096)
		for {
			n, err := r.Read(buf)
			if n == 0 || err == io.EOF {
				break
			} else if n < 0 || err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
				break
			}
			fmt.Printf("%v", buf[:n])
		}
	},
}

func init() {
	fileCreateCmd.PersistentFlags().StringVarP(&fileCreateParams.File, "file", "f", "", "Name of the JSON Lines file to be uploaded.\nIf the purpose is set to \"fine-tune\", each line is a JSON record with \"prompt\" and \"completion\" fields representing your training examples.")
	fileCreateCmd.PersistentFlags().StringVarP(&fileCreateParams.Purpose, "purpose", "p", "", "The intended purpose of the uploaded documents.\nUse \"fine-tune\" for Fine-tuning. This allows us to validate the format of the uploaded file.")
	fileCreateCmd.MarkPersistentFlagRequired("file")
	fileCreateCmd.MarkPersistentFlagRequired("purpose")

	fileDeleteCmd.PersistentFlags().StringVarP(&fileDeleteParams.FileID, "file-id", "f", "", "The ID of the file to use for this request")
	fileDeleteCmd.MarkPersistentFlagRequired("file-id")

	fileRetrieveCmd.PersistentFlags().StringVarP(&fileRetrieveParams.FileID, "file-id", "f", "", "The ID of the file to use for this request")
	fileRetrieveCmd.MarkPersistentFlagRequired("file-id")

	fileDownloadCmd.PersistentFlags().StringVarP(&fileDownloadParams.FileID, "file-id", "f", "", "The ID of the file to use for this request")
	fileDownloadCmd.MarkPersistentFlagRequired("file-id")

	FileCmd.AddCommand(fileCreateCmd, fileDeleteCmd, fileRetrieveCmd, fileDownloadCmd)
}
