package model

import (
	"reflect"
	"regexp"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New(validator.WithRequiredStructEnabled())

	validate.RegisterValidation("brand_name", func(fl validator.FieldLevel) bool {
		if fl.Field().Kind() != reflect.String {
			return false
		}

		// Must be one of
		// - word character
		// - space
		// - dash
		// - ampersand
		match, err := regexp.MatchString("^[\\w\\s\\-&]+$", fl.Field().String())
		if err != nil {
			return false
		}

		return match
	})

	validate.RegisterValidation("currency", func(fl validator.FieldLevel) bool {
		if fl.Field().Kind() != reflect.String {
			return false
		}

		// Must have leading digits and only 2 decimal digits
		match, err := regexp.MatchString("^\\d+\\.\\d{2}$", fl.Field().String())
		if err != nil {
			return false
		}

		return match
	})
}
