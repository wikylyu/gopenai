package files

import (
	"encoding/json"
	"io"

	"github.com/wikylyu/gopenai/client"
)

func NewClient(c *client.Client) *FileClient {
	return &FileClient{c: c}
}

/*
 * Returns a list of files that belong to the user's organization.
 */
func (c *FileClient) List() (*ListResponse, error) {
	data, err := c.c.DoJson("GET", "/files", nil)
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
 * Upload a file that contains document(s) to be used across various endpoints/features.
 * Currently, the size of all the files uploaded by one organization can be up to 1 GB.
 * Please contact us if you need to increase the storage limit.
 */
func (c *FileClient) Create(req *CreateRequest) (*File, error) {
	data, err := c.c.DoForm("POST", "/files", req)
	if err != nil {
		return nil, err
	}
	var resp File
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

/*
 * Delete a file.
 */
func (c *FileClient) Delete(id string) (*DeleteResponse, error) {
	data, err := c.c.DoJson("DELETE", "/files/"+id, nil)
	if err != nil {
		return nil, err
	}
	var resp DeleteResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

/*
 * Returns information about a specific file.
 */
func (c *FileClient) Retrieve(id string) (*File, error) {
	data, err := c.c.DoJson("GET", "/files/"+id, nil)
	if err != nil {
		return nil, err
	}
	var resp File
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

/*
 * Returns the contents of the specified file
 */
func (c *FileClient) Download(id string) (io.ReadCloser, error) {
	r, err := c.c.Download("/files/" + id + "/content")
	if err != nil {
		return nil, err
	}
	return r, nil
}
