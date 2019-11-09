package gons3

import "errors"

// ErrorWrapper wraps two errors
type ErrorWrapper struct {
	Current error
	Next    error
}

// Error message
func (e ErrorWrapper) Error() string {
	return e.Current.Error() + ": " + e.Next.Error()
}

// Unwrap error
func (e ErrorWrapper) Unwrap() error {
	return e.Next
}

// Is error
func (e ErrorWrapper) Is(err error) bool {
	return errors.Is(err, e.Current)
}

// As error
func (e ErrorWrapper) As(err error, target interface{}) bool {
	return errors.As(err, e.Current)
}

// WrapErrors w
func WrapErrors(errs ...error) error {
	if len(errs) == 0 {
		return nil
	} else if len(errs) == 1 {
		return errs[0]
	}

	last := errs[len(errs)-1]
	for i := len(errs) - 2; i >= 0; i-- {
		last = ErrorWrapper{Current: errs[i], Next: last}
	}

	return last
}
