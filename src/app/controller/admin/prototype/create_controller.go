package controller

import (
	dtos "github.com/pudthaiiii/golang-cms/src/app/controller/admin/dtos/request/prototype"
	"github.com/pudthaiiii/golang-cms/src/app/validator"

	"github.com/gofiber/fiber/v2"
)

func (p prototypeController) Create(c *fiber.Ctx) error {
	var data dtos.PrototypeCreateRequest

	if err := validator.ValidateRequest(c, &data); err != nil {
		return err
	}

	return nil
}
