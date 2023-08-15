package errorPackage

import (
	"fmt"
)

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("API error %d: %s", e.Code, e.Message)
}

//pre-defined errors
var (
	ErrBadRequest = &APIError{Code: 400, Message: "Bad Request"}
	ErrNotFound   = &APIError{Code: 404, Message: "Not Found"}
)
