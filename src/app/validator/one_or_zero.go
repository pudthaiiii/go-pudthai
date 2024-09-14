package validator

import "github.com/go-playground/validator/v10"

func init() {
	validate.RegisterValidation("oneOrZero", validateOneOrZero)
}

func validateOneOrZero(fl validator.FieldLevel) bool {
	value := fl.Field().Int()
	return value == 0 || value == 1
}
