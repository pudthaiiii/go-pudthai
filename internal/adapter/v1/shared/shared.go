package shared

import (
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

func Success(c *fiber.Ctx, data, pagination interface{}) error {
	statusCode := getStatus(c)

	response := successResponse{
		Status: Status{
			Code:    statusCode,
			Message: "OK",
		},
		Data:       data,
		Pagination: pagination,
	}

	return c.Status(statusCode).JSON(response)
}

func getStatus(c *fiber.Ctx) int {
	if c.Route().Name == "Create" {
		return fiber.StatusCreated
	}

	return fiber.StatusOK
}
