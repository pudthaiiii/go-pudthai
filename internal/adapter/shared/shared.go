package shared

import (
	"go-pudthai/internal/model/technical"

	"github.com/gofiber/fiber/v2"
)

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type successResponse struct {
	Data       interface{} `json:"data,omitempty"`
	Pagination interface{} `json:"pagination,omitempty"`
	Status     Status      `json:"status"`
}

func Success(c *fiber.Ctx, data, pagination interface{}, statusCode ...int) error {
	defaultStatusCode := 200

	if len(statusCode) > 0 {
		defaultStatusCode = statusCode[0]
	}

	response := successResponse{
		Status: Status{
			Code:    defaultStatusCode,
			Message: technical.HttpStatusMessages[defaultStatusCode],
		},
		Data:       data,
		Pagination: pagination,
	}

	return c.Status(defaultStatusCode).JSON(response)
}
