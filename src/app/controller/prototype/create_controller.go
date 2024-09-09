package controller

import (
	request "github.com/pudthaiiii/golang-cms/src/app/requests/prototype"
	"github.com/pudthaiiii/golang-cms/src/app/validator"

	"github.com/gofiber/fiber/v2"
)

func (p prototypeController) Create(c *fiber.Ctx) error {
	var data request.PrototypeCreateRequest

	if err := validator.ParseAndValidateRequest(c, &data); err != nil {
		return err
	}

	return nil
}
