package controllers

import (
	ApiResource "github.com/pudthaiiii/go-ibooking/src/resource"

	"github.com/gofiber/fiber/v2"
)

func (p prototypeController) Paginate(c *fiber.Ctx) error {
	data, err := p.prototypeService.Paginate(c.Context())
	if err != nil {
		return err
	}

	return ApiResource.SuccessResponse(c, data, nil, nil)
}
