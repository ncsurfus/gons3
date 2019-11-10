package gons3

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// WrappedError implements a linked list of "Wrapped" errors.
// This supports errors.Unwrap, errors.Is, and errors.As
type WrappedError struct {
	Current error
	Next    error
}

// Error returns the complete error message of linked errors.
func (e WrappedError) Error() string {
	return e.Current.Error() + ": " + e.Next.Error()
}

// Unwrap error returns the wrapped error.
func (e WrappedError) Unwrap() error {
	return e.Next
}

// Is implements "errors.Is" for the current and wrapped errors.
func (e WrappedError) Is(err error) bool {
	return errors.Is(err, e.Current)
}

// As implements "errors.As" for the current and wrapped errors.
func (e WrappedError) As(err error, target interface{}) bool {
	return errors.As(err, e.Current)
}

// Wrap wraps the errors so the first wraps the second, and the second wraps the third, etc..
func Wrap(errs ...error) error {
	if len(errs) == 0 {
		return nil
	} else if len(errs) == 1 {
		return errs[0]
	}

	last := errs[len(errs)-1]
	for i := len(errs) - 2; i >= 0; i-- {
		last = WrappedError{Current: errs[i], Next: last}
	}

	return last
}

// ServerError represents a GNS3 server error and message
type ServerError struct {
	msg  string
	code int
}

// GetStatusCode gets the status code from the server
func (s ServerError) GetStatusCode() int {
	return s.code
}

// Error returns the error message for the ServerError
func (s ServerError) Error() string {
	if s.msg == "" {
		return fmt.Sprintf("status code %v", s.code)
	}
	return fmt.Sprintf("status code %v: %v", s.code, s.msg)
}

func newServerError(resp *http.Response) ServerError {
	// Handle non-JSON error messages
	contentType := resp.Header.Get("Content-type")
	if !strings.Contains(contentType, "application/json") {
		return ServerError{code: resp.StatusCode}
	}

	// Unmarshal response to GNS3 error handler
	// If this errors we'll just show the status code, something has gone wrong!
	// Read body, ignoring error
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ServerError{code: resp.StatusCode}
	}

	// Unmarshal JSON, ignoring error
	j := struct {
		Message string `json:"message"`
	}{}
	if json.Unmarshal(respBody, &j) != nil {
		return ServerError{code: resp.StatusCode}
	}

	// Ignore the schema message which is a rather large escaped JSON payload
	if i := strings.Index(j.Message, " in schema"); i != -1 {
		j.Message = j.Message[:i]
	}

	return ServerError{
		code: resp.StatusCode,
		msg:  j.Message,
	}
}
