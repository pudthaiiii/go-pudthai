package controller

import (
	"workshop/src/app/services"

	"github.com/gofiber/fiber/v2"
)

type prototypeController struct {
	prototypeService services.PrototypeInteractor
}

type PrototypeController interface {
	Paginate(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
}

func NewPrototypeController(prototypeService services.PrototypeInteractor) PrototypeController {
	return &prototypeController{
		prototypeService,
	}
}
