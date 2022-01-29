package httputil

import "net/http"

var (
	ErrUnauthorized = NewError(http.StatusUnauthorized, "user.not_authorized")
)

// Error represents a handler error. It provides methods for a HTTP status
// code and embeds the built-in error interface.
type Error interface {
	error
	Status() int
}

// StatusError represents error with http status code.
type StatusError struct {
	Code     int      `json:"code"`
	Message  string   `json:"message,omitempty"`
	Messages []string `json:"messages,omitempty"`
}

// Error allows StatusError to satisfy the error interface.
func (se StatusError) Error() string {
	return se.Message
}

// Status returns our HTTP status code.
func (se StatusError) Status() int {
	return se.Code
}

// NewError creates new error instance.
func NewError(code int, messages ...string) Error {
	statusErr := StatusError{Code: code}

	if len(messages) == 1 {
		statusErr.Message = messages[0]
	} else {
		statusErr.Messages = messages
	}

	return statusErr
}
