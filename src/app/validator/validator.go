package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/iancoleman/strcase"
)

var validate = validator.New()

type (
	errorResponse struct {
		Error       bool
		FailedField string
		Tag         string
	}

	xValidator struct {
		validator *validator.Validate
	}
)

func (v xValidator) inValidate(data interface{}) []errorResponse {
	validationErrors := []errorResponse{}

	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			var elem errorResponse

			elem.FailedField = err.Field()
			elem.Tag = err.Tag()
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}

func ValidateRequest[T any](c *fiber.Ctx, data *T) error {
	myValidator := &xValidator{
		validator: validate,
	}

	var err error

	switch c.Method() {
	case "GET":
		err = c.QueryParser(data)
	default:
		err = c.BodyParser(data)
	}

	if err != nil {
		errorMsg := "Cannot parse query parameters"

		if c.Method() != "GET" {
			errorMsg = "Cannot parse JSON"
		}

		return handleValidationError(c, errorMsg)
	}

	if errs := myValidator.inValidate(data); len(errs) > 0 && errs[0].Error {
		errorMap := make(map[string][]string)

		for _, err := range errs {
			failedField := strcase.ToLowerCamel(err.FailedField)

			if _, ok := errorMap[failedField]; !ok {
				errorMap[failedField] = []string{}
			}

			errorMsg := fmt.Sprintf("%v failed on the '%s'", failedField, err.Tag)
			errorMap[failedField] = append(errorMap[failedField], errorMsg)
		}

		return handleValidationError(c, errorMap)
	}

	return nil
}

func handleValidationError(c *fiber.Ctx, errors interface{}) error {
	response := map[string]interface{}{
		"status": map[string]interface{}{
			"code":    422,
			"message": "Unprocessable Entity",
		},
		"error": map[string]interface{}{
			"code":    100422,
			"message": "VALIDATE_ERROR",
			"errors":  errors,
		},
	}

	return c.Status(fiber.StatusUnprocessableEntity).JSON(response)
}
