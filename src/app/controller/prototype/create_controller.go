package controller

import (
	requests "workshop/src/app/requests/prototype"
	"workshop/src/app/validator"

	"github.com/gofiber/fiber/v2"
)

func (p prototypeController) Create(c *fiber.Ctx) error {
	var data requests.PrototypeCreateRequest

	if err := validator.ParseAndValidateRequest(c, &data); err != nil {
		return err
	}

	return nil
}
