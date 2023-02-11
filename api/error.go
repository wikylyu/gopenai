package api

import "fmt"

type Error struct {
	Message string      `json:"message"`
	Type    string      `json:"type"`
	Param   interface{} `json:"param"`
	Code    string      `json:"code"`
}

func (e *Error) Error() string {
	return fmt.Sprintf(`[%s]:%s`, e.Type, e.Message)
}

type ErrorResponse struct {
	Error Error `json:"error"`
}
