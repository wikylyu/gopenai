package finetunes

import (
	"encoding/json"

	"github.com/wikylyu/gopenai/client"
)

func NewClient(c *client.Client) *FineTuneClient {
	return &FineTuneClient{
		c: c,
	}
}

/*
 * Creates a job that fine-tunes a specified model from a given dataset.
 * Response includes details of the enqueued job including job status and the name of the fine-tuned models once complete.
 */
func (c *FineTuneClient) Create(req *CreateRequest) (*FineTune, error) {
	data, err := c.c.DoJson("POST", "/fine-tunes", req)
	if err != nil {
		return nil, err
	}
	var resp FineTune
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

/*
 * List your organization's fine-tuning jobs
 */
func (c *FineTuneClient) List() (*ListResponse, error) {
	data, err := c.c.DoJson("GET", "/fine-tunes", nil)
	if err != nil {
		return nil, err
	}
	var resp ListResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

/*
 * Gets info about the fine-tune job.
 */
func (c *FineTuneClient) Retrieve(id string) (*FineTune, error) {
	data, err := c.c.DoJson("GET", "/fine-tunes/"+id, nil)
	if err != nil {
		return nil, err
	}
	var resp FineTune
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

/*
 * Immediately cancel a fine-tune job.
 */
func (c *FineTuneClient) Cancel(id string) (*FineTune, error) {
	data, err := c.c.DoJson("POST", "/fine-tunes/"+id+"/cancel", nil)
	if err != nil {
		return nil, err
	}
	var resp FineTune
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

/*
 * Get fine-grained status updates for a fine-tune job.
 */
func (c *FineTuneClient) ListEvents(id string) (*ListEventsResponse, error) {
	data, err := c.c.DoJson("GET", "/fine-tunes/"+id+"/events", nil)
	if err != nil {
		return nil, err
	}
	var resp ListEventsResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
