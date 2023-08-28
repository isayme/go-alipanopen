package alipanopen

import "fmt"

type ErrorResponse struct {
	StatusCode int    `json:"-"`
	Code       string `json:"code"`
	Message    string `json:"message"`
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("%s(%d): %s", e.Code, e.StatusCode, e.Message)
}
