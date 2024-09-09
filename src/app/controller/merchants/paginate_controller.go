package controller

import (
	"github.com/gofiber/fiber/v2"
)

// GetAll godoc
// @Tags Plant
// @Summary get plant
// @Description get plant
// @ID get-plant
// @Router /prototype [get]
func (r merchantsController) GetAll(c *fiber.Ctx) error {
	data, err := r.prototypeInteractor.Paginate(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Hello, World!",
		"user":    data,
	})
}
