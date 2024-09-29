package throw

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func Error(code int, err error, statusCode ...int) error {
	defaultStatusCode := 422

	if len(statusCode) > 0 {
		defaultStatusCode = statusCode[0]
	}

	response := map[string]interface{}{
		"code":    code,
		"message": ErrorCodes[code],
	}

	if err != nil {
		response["error"] = err.Error()
	}

	responseJSON, _ := json.Marshal(response)
	return fiber.NewError(defaultStatusCode, string(responseJSON))
}
