package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/wikylyu/gopenai/finetunes"
)

var FineTuneCmd = &cobra.Command{
	Use:   "fine-tune",
	Short: "Manage fine-tuning jobs to tailor a model to your specific training data.",
	Long:  "Manage fine-tuning jobs to tailor a model to your specific training data.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var fineTuneCreateParams struct {
	TrainingFile                 string
	ValidationFile               string
	Model                        string
	NEpochs                      int64
	BatchSize                    int64
	LearningRateMultiplier       float64
	PromptLossWeight             float64
	ComputeClassificationMetrics bool
	ClassificationNClasses       int64
	ClassificationPositiveClass  string
	ClassificationBetas          []float64
	Suffix                       string
}

var fineTuneCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a job that fine-tunes a specified model from a given dataset.",
	Long:  "Creates a job that fine-tunes a specified model from a given dataset.\nResponse includes details of the enqueued job including job status and the name of the fine-tuned models once complete.",
	Run: func(cmd *cobra.Command, args []string) {
		trainingFile := fineTuneCreateParams.TrainingFile
		validationFile := fineTuneCreateParams.ValidationFile
		model := fineTuneCreateParams.Model
		nEpochs := fineTuneCreateParams.NEpochs
		batchSize := fineTuneCreateParams.BatchSize
		learningRateMultiplier := fineTuneCreateParams.LearningRateMultiplier
		promptLossWeight := fineTuneCreateParams.PromptLossWeight
		computeClassificationMetrics := fineTuneCreateParams.ComputeClassificationMetrics
		classificationNClasses := fineTuneCreateParams.ClassificationNClasses
		classificationPositiveClass := fineTuneCreateParams.ClassificationPositiveClass
		classificationBetas := fineTuneCreateParams.ClassificationBetas
		suffix := fineTuneCreateParams.Suffix

		r, err := openai.FineTune.Create(&finetunes.CreateRequest{
			TrainingFile:                 trainingFile,
			ValidationFile:               validationFile,
			Model:                        model,
			NEpochs:                      nEpochs,
			BatchSize:                    batchSize,
			LearningRateMultiplier:       learningRateMultiplier,
			PromptLossWeight:             promptLossWeight,
			ComputeClassificationMetrics: computeClassificationMetrics,
			ClassificationNClasses:       classificationNClasses,
			ClassificationPositiveClass:  classificationPositiveClass,
			ClassificationBetas:          classificationBetas,
			Suffix:                       suffix,
		})
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			return
		}
		printJson(r)
	},
}

var fineTuneListCmd = &cobra.Command{
	Use:   "list",
	Short: "List your organization's fine-tuning jobs.",
	Long:  "List your organization's fine-tuning jobs.",
	Run: func(cmd *cobra.Command, args []string) {
		r, err := openai.FineTune.List()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			return
		}
		printJson(r)
	},
}

var fineTuneRetrieveParams struct {
	FineTuneID string
}

var fineTuneRetrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Gets info about the fine-tune job.",
	Long:  "Gets info about the fine-tune job.",
	Run: func(cmd *cobra.Command, args []string) {
		fineTuneID := fineTuneRetrieveParams.FineTuneID
		r, err := openai.FineTune.Retrieve(fineTuneID)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			return
		}
		printJson(r)
	},
}

var fineTuneCancelParams struct {
	FineTuneID string
}

var fineTuneCancelCmd = &cobra.Command{
	Use:   "cancel",
	Short: "Immediately cancel a fine-tune job.",
	Long:  "Immediately cancel a fine-tune job.",
	Run: func(cmd *cobra.Command, args []string) {
		fineTuneID := fineTuneCancelParams.FineTuneID
		r, err := openai.FineTune.Cancel(fineTuneID)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			return
		}
		printJson(r)
	},
}

var fineTuneListEventsParams struct {
	FineTuneID string
}

var fineTuneListEventsCmd = &cobra.Command{
	Use:   "list-events",
	Short: "Get fine-grained status updates for a fine-tune job.",
	Long:  "Get fine-grained status updates for a fine-tune job.",
	Run: func(cmd *cobra.Command, args []string) {
		fineTuneID := fineTuneListEventsParams.FineTuneID
		r, err := openai.FineTune.ListEvents(fineTuneID)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			return
		}
		printJson(r)
	},
}

func init() {
	fineTuneCreateCmd.PersistentFlags().StringVarP(&fineTuneCreateParams.TrainingFile, "training-file", "t", "", "The ID of an uploaded file that contains training data.\nSee upload file for how to upload a file.")
	fineTuneCreateCmd.PersistentFlags().StringVarP(&fineTuneCreateParams.ValidationFile, "validation-file", "v", "", "The ID of an uploaded file that contains validation data.\nIf you provide this file, the data is used to generate validation metrics periodically during fine-tuning. These metrics can be viewed in the fine-tuning results file. Your train and validation data should be mutually exclusive.\nYour dataset must be formatted as a JSONL file, where each validation example is a JSON object with the keys \"prompt\" and \"completion\". Additionally, you must upload your file with the purpose fine-tune.")
	fineTuneCreateCmd.PersistentFlags().StringVarP(&fineTuneCreateParams.Model, "model", "m", "curie", "The name of the base model to fine-tune. You can select one of \"ada\", \"babbage\", \"curie\", \"davinci\", or a fine-tuned model created after 2022-04-21. To learn more about these models.")
	fineTuneCreateCmd.PersistentFlags().Int64VarP(&fineTuneCreateParams.NEpochs, "n-epochs", "n", 4, "The number of epochs to train the model for. An epoch refers to one full cycle through the training dataset.")
	fineTuneCreateCmd.PersistentFlags().Int64VarP(&fineTuneCreateParams.BatchSize, "batch-size", "b", 0, "The batch size to use for training. The batch size is the number of training examples used to train a single forward and backward pass.\nBy default, the batch size will be dynamically configured to be ~0.2% of the number of examples in the training set, capped at 256 - in general, we've found that larger batch sizes tend to work better for larger datasets.")
	fineTuneCreateCmd.PersistentFlags().Float64VarP(&fineTuneCreateParams.LearningRateMultiplier, "learning-rate-multiplier", "l", 0, "The learning rate multiplier to use for training. The fine-tuning learning rate is the original learning rate used for pretraining multiplied by this value.\nBy default, the learning rate multiplier is the 0.05, 0.1, or 0.2 depending on final batch_size (larger learning rates tend to perform better with larger batch sizes). We recommend experimenting with values in the range 0.02 to 0.2 to see what produces the best results.")
	fineTuneCreateCmd.PersistentFlags().Float64VarP(&fineTuneCreateParams.PromptLossWeight, "prompt-loss-weight", "w", 0.01, "The weight to use for loss on the prompt tokens. This controls how much the model tries to learn to generate the prompt (as compared to the completion which always has a weight of 1.0), and can add a stabilizing effect to training when completions are short.\nIf prompts are extremely long (relative to completions), it may make sense to reduce this weight so as to avoid over-prioritizing learning the prompt.")
	fineTuneCreateCmd.PersistentFlags().BoolVarP(&fineTuneCreateParams.ComputeClassificationMetrics, "compute-classification-metrics", "c", false, "If set, we calculate classification-specific metrics such as accuracy and F-1 score using the validation set at the end of every epoch. These metrics can be viewed in the results file.\nIn order to compute classification metrics, you must provide a validation_file. Additionally, you must specify classification_n_classes for multiclass classification or classification_positive_class for binary classification.")
	fineTuneCreateCmd.PersistentFlags().Int64VarP(&fineTuneCreateParams.ClassificationNClasses, "classification-n-classes", "", 0, "The number of classes in a classification task.\nThis parameter is required for multiclass classification.")
	fineTuneCreateCmd.PersistentFlags().StringVarP(&fineTuneCreateParams.ClassificationPositiveClass, "classification-positive-class", "", "", "The positive class in binary classification.\nThis parameter is needed to generate precision, recall, and F1 metrics when doing binary classification.")
	fineTuneCreateCmd.PersistentFlags().Float64SliceVarP(&fineTuneCreateParams.ClassificationBetas, "classification-betas", "", nil, "If this is provided, we calculate F-beta scores at the specified beta values. The F-beta score is a generalization of F-1 score. This is only used for binary classification.\nWith a beta of 1 (i.e. the F-1 score), precision and recall are given the same weight. A larger beta score puts more weight on recall and less on precision. A smaller beta score puts more weight on precision and less on recall.")
	fineTuneCreateCmd.PersistentFlags().StringVarP(&fineTuneCreateParams.Suffix, "suffix", "s", "", "A string of up to 40 characters that will be added to your fine-tuned model name.\nFor example, a suffix of \"custom-model-name\" would produce a model name like ada:ft-your-org:custom-model-name-2022-02-15-04-21-04.")
	fineTuneCreateCmd.MarkPersistentFlagRequired("training-file")

	fineTuneRetrieveCmd.PersistentFlags().StringVarP(&fineTuneRetrieveParams.FineTuneID, "fine-tune-id", "f", "", "The ID of the fine-tune job.")
	fineTuneRetrieveCmd.MarkPersistentFlagRequired("fine-tune-id")

	fineTuneCancelCmd.PersistentFlags().StringVarP(&fineTuneCancelParams.FineTuneID, "fine-tune-id", "f", "", "The ID of the fine-tune job.")
	fineTuneCancelCmd.MarkPersistentFlagRequired("fine-tune-id")

	fineTuneListEventsCmd.PersistentFlags().StringVarP(&fineTuneListEventsParams.FineTuneID, "fine-tune-id", "f", "", "The ID of the fine-tune job.")
	fineTuneListEventsCmd.MarkPersistentFlagRequired("fine-tune-id")

	FineTuneCmd.AddCommand(fineTuneCreateCmd, fineTuneListCmd, fineTuneRetrieveCmd, fineTuneCancelCmd, fineTuneListEventsCmd)
}
