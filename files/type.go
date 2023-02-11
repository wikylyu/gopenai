package files

import (
	"os"

	"github.com/wikylyu/gopenai/api"
)

type FileClient struct {
	c *api.Client
}

type File struct {
	ID            string                 `json:"id"`
	Object        string                 `json:"object"`
	Bytes         int64                  `json:"bytes"`
	CreatedAt     int64                  `json:"created_at"`
	Filename      string                 `json:"filename"`
	Purpose       string                 `json:"purpose"`
	Status        string                 `json:"status"`
	StatusDetails map[string]interface{} `json:"status_details"`
}

type ListResponse struct {
	Object string  `json:"object"`
	Data   []*File `json:"data"`
}

type CreateRequest struct {
	File    *os.File `json:"file" form:"file"`
	Purpose string   `json:"purpose" form:"purpose"`
}

type DeleteResponse struct {
	ID      string `json:"id"`
	Object  string `json:"file"`
	Deleted bool   `json:"deleted"`
}
