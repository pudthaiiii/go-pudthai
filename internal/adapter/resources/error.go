package resources

import (
	"encoding/json"
	"errors"
	"go-ibooking/src/utils"
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
	statusCode := http.StatusInternalServerError
	errorCode := 0
	errorMessage := "Internal Server Error"
	var exception []string

	var e *fiber.Error

	if errors.As(err, &e) {
		statusCode = e.Code
		errorCode = e.Code
		errorMessage = e.Message

		if strings.Contains(e.Error(), "VALIDATE_ERROR") {
			statusCode = http.StatusUnprocessableEntity
			errorCode = 900422
			errorMessage = "VALIDATE_ERROR"

			var validationError validationError
			json.Unmarshal([]byte(e.Error()), &validationError)

			for _, errs := range validationError.Errors {
				exception = append(exception, errs...)
			}
		}
	} else if err != nil {
		throwException := utils.FilterThrowExceptions(err.Error())

		if len(throwException) > 0 {
			statusCode = http.StatusUnprocessableEntity
			if len(throwException) > 1 && throwException[1] != "" {
				errorCode, _ = strconv.Atoi(throwException[1])
			}

			if len(throwException) > 2 && throwException[2] != "" {
				errorMessage = throwException[2]
			}

			if len(throwException) > 3 && throwException[3] != "" {
				exception = append(exception, throwException[3])
			}
		}
	}

	response := fiber.Map{
		"message": errorMessage,
		"code":    errorCode,
	}

	if boolTrue, _ := strconv.ParseBool(os.Getenv("EXCEPTION_LOG_ENABLED")); boolTrue && len(exception) != 0 {
		response["errors"] = exception
	}

	return c.Status(statusCode).JSON(response)
}
