package ApiResource

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type validationError struct {
	Errors  map[string][]string `json:"errors"`
	Message string              `json:"message"`
}

func ErrorHandler(c *fiber.Ctx, err error) error {
	errorCode := http.StatusInternalServerError
	errorMessage := "Internal Server Error"
	var exception []string

	var e *fiber.Error

	if errors.As(err, &e) {
		errorCode = e.Code
		errorMessage = e.Message

		if strings.Contains(e.Error(), "VALIDATE_ERROR") {
			errorCode = http.StatusUnprocessableEntity
			errorMessage = "Validation Failed"

			var validationError validationError
			json.Unmarshal([]byte(e.Error()), &validationError)

			for _, errs := range validationError.Errors {
				exception = append(exception, errs...)
			}
		}
	} else if err != nil {
		switch {
		case strings.Contains(err.Error(), `does not exist`):
			errorCode = http.StatusInternalServerError
			errorMessage = "Something went wrong"
			exception = append(exception, err.Error())

		case strings.Contains(err.Error(), "Cannot GET"):
			errorCode = http.StatusNotFound
			errorMessage = "404 Not Found"

		case strings.Contains(err.Error(), "Cannot POST"), strings.Contains(err.Error(), "Cannot DELETE"),
			strings.Contains(err.Error(), "Cannot PATCH"), strings.Contains(err.Error(), "Cannot PUT"):
			errorCode = http.StatusMethodNotAllowed
			errorMessage = "405 Method Not Allowed"

		case strings.Contains(err.Error(), "violates foreign key constraint"):
			errorCode = http.StatusBadRequest
			errorMessage = "Foreign Key Constraint Violation"
			exception = append(exception, err.Error())
		}
	}

	response := fiber.Map{
		"message": errorMessage,
		"code":    errorCode,
	}

	if boolTrue, _ := strconv.ParseBool(os.Getenv("EXCEPTION_LOG_ENABLED")); boolTrue && len(exception) != 0 {
		response["errors"] = exception
	}

	return c.Status(errorCode).JSON(response)
}
