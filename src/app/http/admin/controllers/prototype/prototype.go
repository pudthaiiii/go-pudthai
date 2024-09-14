package controllers

import (
	service "go-ibooking/src/app/services/admin"

	"github.com/gofiber/fiber/v2"
)

type PrototypeController interface {
	Paginate(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
}

type prototypeController struct {
	prototypeService service.PrototypeService
}

func NewPrototypeController(prototypeService service.PrototypeService) PrototypeController {
	return &prototypeController{
		prototypeService,
	}
}
