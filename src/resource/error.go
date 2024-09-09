package ApiResource

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"workshop/src/utils"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	errorCode := http.StatusInternalServerError
	errorMessage := "Internal Server Error"
	var exception []string

	// code := fiber.StatusInternalServerError
	var e *fiber.Error
	// if errors.As(err, &e) {
	// 	code = e.Code
	// }

	fmt.Println(errors.As(err, &e), e.Code)

	if err != nil {
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
		}

		response := fiber.Map{
			"message": errorMessage,
			"code":    errorCode,
		}

		if boolTrue, _ := strconv.ParseBool(utils.RequireEnv("EXCEPTION_LOG_ENABLED", "false")); boolTrue && len(exception) != 0 {
			response["errors"] = exception
		}

		return c.Status(errorCode).JSON(response)
	}

	return nil
}
