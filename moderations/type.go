package moderations

import "github.com/wikylyu/gopenai/client"

type ModerationClient struct {
	c *client.Client
}

type CreateRequest struct {
	Input interface{} `json:"input"`
	Model string      `json:"model,omitempty"`
}

type Categories struct {
	Hate            bool `json:"hate"`
	HateThreatening bool `json:"hate/threatening"`
	SelfHarm        bool `json:"self-harm"`
	Sexual          bool `json:"sexual"`
	SexualMinors    bool `json:"sexual/minors"`
	Violence        bool `json:"violence"`
	ViolenceGraphic bool `json:"violence/graphic"`
}

type CategoryScores struct {
	Hate            float64 `json:"hate"`
	HateThreatening float64 `json:"hate/threatening"`
	SelfHarm        float64 `json:"self-harm"`
	Sexual          float64 `json:"sexual"`
	SexualMinors    float64 `json:"sexual/minors"`
	Violence        float64 `json:"violence"`
	ViolenceGraphic float64 `json:"violence/graphic"`
}

type Moderation struct {
	Categories     Categories     `json:"categories"`
	CategoryScores CategoryScores `json:"category_scores"`
	Flagged        bool           `json:"flagged"`
}

type CreateResponse struct {
	ID      string        `json:"id"`
	Model   string        `json:"model"`
	Results []*Moderation `json:"results"`
}
