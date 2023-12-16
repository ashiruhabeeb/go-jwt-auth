package validator

import "github.com/go-playground/validator/v10"

func ValidateInput(i interface{}) error {
	val := validator.New()
	
	return val.Struct(i)
}
