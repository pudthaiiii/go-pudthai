package controllers

import (
	dtos "github.com/pudthaiiii/go-ibooking/src/app/http/admin/dtos/request/prototype"
	"github.com/pudthaiiii/go-ibooking/src/app/validator"

	"github.com/gofiber/fiber/v2"
)

func (p prototypeController) Create(c *fiber.Ctx) error {
	var data dtos.PrototypeCreateRequest

	if err := validator.Validate(c, &data); err != nil {
		return err
	}

	return nil
}
