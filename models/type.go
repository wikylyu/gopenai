package models

import "github.com/wikylyu/gopenai/client"

type Permission struct {
	ID                 string      `json:"id"`
	Object             string      `json:"object"`
	Created            int64       `json:"created"`
	AllowCreateEngine  bool        `json:"allow_create_engine"`
	AllowSampling      bool        `json:"allow_sampling"`
	AllowLogprobs      bool        `json:"allow_logprobs"`
	AllowSearchIndices bool        `json:"allow_search_indices"`
	AllowView          bool        `json:"allow_view"`
	AllowFineTuning    bool        `json:"allow_fine_tuning"`
	Organization       string      `json:"organization"`
	IsBlocking         bool        `json:"is_blocking"`
	Group              interface{} `json:"group"` // FIXME: don't know what it is.
}

type Model struct {
	ID         string        `json:"id"`
	Object     string        `json:"object"`
	OwnedBy    string        `json:"owned_by"`
	Created    int64         `json:"created"`
	Root       string        `json:"root"`
	Parent     interface{}   `json:"parent"` //FIXME: don't know what it is
	Permission []*Permission `json:"permission"`
}

type ListResponse struct {
	Data   []*Model `json:"data"`
	Object string   `json:"object"`
}

type DeleteResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Deleted bool   `json:"deleted"`
}

type ModelClient struct {
	c *client.Client
}

/*
 * Some frequently-used model name
 */
const (
	ModelAda                 = "ada"
	ModelBabbage             = "babbage"
	ModelDavinci             = "davinci"
	ModelTextDavinci003      = "text-davinci-003"
	ModelTextDavinci002      = "text-davinci-002"
	ModelTextDavinci001      = "text-davinci-001"
	ModelCurieInstructBeta   = "curie-instruct-beta"
	ModelCodeCushman001      = "code-cushman-001"
	ModelTextAda001          = "text-ada-001"
	ModelTextCurie001        = "text-curie-001"
	ModelCodeDavinci002      = "code-davinci-002"
	ModelDavinciInstructBeta = "davinci-instruct-beta"
	ModelTextBabbage001      = "text-babbage-001"
	ModelCurie               = "curie"

	ModelTextDavinciEdit001 = "text-davinci-edit-001"
	ModelCodeDavinciEdit001 = "code-davinci-edit-001"

	ModelTextEmbeddingAda002 = "text-embedding-ada-002"
)
