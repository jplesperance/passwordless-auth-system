package utils

import (
	"fmt"
)

type HttpResponse struct {
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	StatusCode int         `json:"statusCode"`
}

func (e HttpResponse) Error() string {
	return fmt.Sprintf("description: %s", e.Message)
}

func NewHttpResponse(statusCode int, message string, data interface{}) HttpResponse {
	return HttpResponse{
		Message:    message,
		Data:       data,
		StatusCode: statusCode,
	}
}
