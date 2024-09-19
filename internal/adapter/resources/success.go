package resources

import (
	"fmt"

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

func SuccessResponse(c *fiber.Ctx, data, pagination interface{}) error {
	statusCode := fiber.StatusOK

	if c.Route().Name == "Create" {
		fmt.Println("Create")
		statusCode = fiber.StatusCreated
	}

	return c.
		Status(statusCode).
		JSON(successResponse{
			Status: Status{
				Code:    statusCode,
				Message: "OK",
			},
			Data:       data,
			Pagination: pagination,
		})
}
