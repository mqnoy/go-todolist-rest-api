package cvalidator

import (
	"strings"
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
		Validator.RegisterValidation("normalize", normalizeString)
	})
}

func normalizeString(fl validator.FieldLevel) bool {
	str := fl.Field().String()
	strLower := strings.ToLower(str)
	fl.Field().SetString(strLower)
	return true
}
