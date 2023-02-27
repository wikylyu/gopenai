package finetunes

import (
	"github.com/wikylyu/gopenai/api"
	"github.com/wikylyu/gopenai/files"
)

type FineTuneClient struct {
	c *api.Client
}

type CreateRequest struct {
	TrainingFile                 string    `json:"training_file"`
	ValidationFile               string    `json:"validation_file,omitempty"`
	Model                        string    `json:"model,omitempty"`
	NEpochs                      int64     `json:"n_epochs,omitempty"`
	BatchSize                    int64     `json:"batch_size,omitempty"`
	LearningRateMultiplier       float64   `json:"learning_rate_multiplier,omitempty"`
	PromptLossWeight             float64   `json:"prompt_loss_weight,omitempty"`
	ComputeClassificationMetrics bool      `json:"compute_classification_metrics,omitempty"`
	ClassificationNClasses       int64     `json:"classification_n_classes,omitempty"`
	ClassificationPositiveClass  string    `json:"classification_positive_class,omitempty"`
	ClassificationBetas          []float64 `json:"classification_betas,omitempty"`
	Suffix                       string    `json:"suffix,omitempty"`
}

type Event struct {
	Object    string `json:"object"`
	CreatedAt int64  `json:"created_at"`
	Level     string `json:"level"`
	Message   string `json:"message"`
}

type FineTune struct {
	ID             string   `json:"id"`
	Object         string   `json:"object"`
	Model          string   `json:"model"`
	CreatedAt      int64    `json:"created_at"`
	UpdatedAt      int64    `json:"updated_at"`
	Events         []*Event `json:"events"`
	FineTunedModel string   `json:"fine_tuned_model"`
	Hyperparams    struct {
		BatchSize              int64   `json:"batch_size"`
		LearningRateMultiplier float64 `json:"learning_rate_multiplier"`
		NEpochs                int64   `json:"n_epochs"`
		PromptLossWeight       float64 `json:"prompt_loss_weight"`
	} `json:"hyperparams"`
	OrganizationID  string        `json:"organization_id"`
	ResultFiles     []*files.File `json:"result_files"`
	Status          string        `json:"status"`
	ValidationFiles []*files.File `json:"validation_files"`
	TrainingFiles   []*files.File `json:"training_files"`
}

type ListResponse struct {
	Object string      `json:"object"`
	Data   []*FineTune `json:"data"`
}

type ListEventsResponse struct {
	Object string   `json:"object"`
	Data   []*Event `json:"data"`
}
