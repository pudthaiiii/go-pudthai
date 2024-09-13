package controllers

import (
	adminService "go-ibooking/src/app/services/admin"

	"github.com/gofiber/fiber/v2"
)

type prototypeController struct {
	prototypeService adminService.PrototypeService
}

type PrototypeController interface {
	Paginate(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
}

func NewPrototypeController(prototypeService adminService.PrototypeService) PrototypeController {
	return &prototypeController{
		prototypeService,
	}
}
