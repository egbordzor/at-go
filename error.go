package atgo

import "fmt"

type APIError struct {
	StatusCode int
	Code       string `json:"code"`
	Message    string `json:"message"`
}

func (err *APIError) Error() string {
	return fmt.Sprintf("%s: %s", err.Code, err.Message)
}
