package cerror

import (
	"errors"
	"fmt"
)

var (
	ErrCantBeEmpty      = errors.New("cant be empty")
	ErrRequiredId       = errors.New("required id")
	ErrResourceNotFound = errors.New("resource not found")
)

type CustomError struct {
	StatusCode int
	Err        error
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("status %d: err %v", e.StatusCode, e.Err)
}

func WrapError(statusCode int, err error) *CustomError {
	return &CustomError{
		StatusCode: statusCode,
		Err:        err,
	}
}
