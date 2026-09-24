package validator

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// ValidateStruct checks if a struct satisfies its validation tags
func ValidateStruct(s interface{}) error {
	return validate.Struct(s)
}
