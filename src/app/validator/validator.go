package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/iancoleman/strcase"
)

type (
	ErrorResponse struct {
		Error       bool
		FailedField string
		Tag         string
		Value       interface{}
	}

	XValidator struct {
		validator *validator.Validate
	}

	GlobalErrorHandlerResp struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}
)

var validate = validator.New()

func (v XValidator) Validate(data interface{}) []ErrorResponse {
	validationErrors := []ErrorResponse{}

	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			// In this case data object is actually holding the User struct
			var elem ErrorResponse
			elem.FailedField = err.Field() // Export struct field name
			elem.Tag = err.Tag()           // Export struct tag
			elem.Value = err.Value()       // Export field value
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}

func ParseAndValidateRequest[T any](c *fiber.Ctx, data *T) error {
	myValidator := &XValidator{
		validator: validate,
	}

	switch c.Method() {
	case "GET":
		if err := c.QueryParser(data); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Cannot parse query parameters")
		}
	case "POST", "PUT", "PATCH", "DELETE":
		if err := c.BodyParser(data); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Cannot parse JSON")
		}
	}

	if errs := myValidator.Validate(data); len(errs) > 0 && errs[0].Error {
		errorMap := make(map[string][]string)
		for _, err := range errs {
			failedField := strcase.ToLowerCamel(err.FailedField)

			if _, ok := errorMap[failedField]; !ok {
				errorMap[failedField] = []string{}
			}

			errorMsg := fmt.Sprintf("%v failed on the '%s'", failedField, err.Tag)
			errorMap[failedField] = append(errorMap[failedField], errorMsg)
		}

		response := map[string]interface{}{
			"status": map[string]interface{}{
				"code":    422,
				"message": "Unprocessable Entity",
			},
			"error": map[string]interface{}{
				"code":    100422,
				"message": "VALIDATE_ERROR",
				"errors":  errorMap,
			},
		}

		return c.Status(fiber.StatusUnprocessableEntity).JSON(response)
	}

	return nil
}
