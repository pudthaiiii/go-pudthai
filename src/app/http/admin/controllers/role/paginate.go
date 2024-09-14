package controllers

import (
	ApiResource "go-ibooking/src/app/resources"

	"github.com/gofiber/fiber/v2"
)

func (s roleController) Paginate(c *fiber.Ctx) error {
	data, err := s.roleService.Paginate(c.Context())
	if err != nil {
		return err
	}

	return ApiResource.SuccessResponse(c, data, nil, nil)
}
