package cvalidator

import (
	"sync"

	"github.com/go-playground/validator/v10"
)

var (
	Validator *validator.Validate
	once      sync.Once
)

const (
	ErrorValidator = "ERROR_VALIDATOR"
)

func init() {
	once.Do(func() {
		Validator = validator.New()

		// Register your custom validator function here
		// ...
	})
}
