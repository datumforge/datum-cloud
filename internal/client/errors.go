package client

import (
	"fmt"
	"strings"
)

// RequestError is a generic error when a request with the client fails
type RequestError struct {
	// StatusCode is the http response code that was returned
	StatusCode int
	// Body of the response
	Body string
}

// Error returns the RequestError in string format
func (e *RequestError) Error() string {
	if e.Body == "" {
		return fmt.Sprintf("unable to process request (status %d)", e.StatusCode)
	}

	return fmt.Sprintf("unable to process request (status %d): %s", e.StatusCode, strings.ToLower(e.Body))
}

// newRequestError returns an error when a datum client request fails
func newRequestError(statusCode int, body string) *RequestError {
	return &RequestError{
		StatusCode: statusCode,
		Body:       body,
	}
}
