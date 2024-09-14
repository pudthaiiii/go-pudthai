package ApiResource

import "github.com/gofiber/fiber/v2"

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type successResponse struct {
	Data   interface{} `json:"data,omitempty"`
	Links  interface{} `json:"links,omitempty"`
	Meta   interface{} `json:"meta,omitempty"`
	Status Status      `json:"status"`
}

func SuccessResponse(c *fiber.Ctx, data, links, meta interface{}) error {
	return c.JSON(successResponse{
		Status: Status{
			Code:    fiber.StatusOK,
			Message: "OK",
		},
		Data:  data,
		Links: links,
		Meta:  meta,
	})
}
