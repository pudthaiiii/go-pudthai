package controller

import (
	"github.com/pudthaiiii/golang-cms/src/app/services"

	"github.com/gofiber/fiber/v2"
)

type merchantsController struct {
	prototypeInteractor services.PrototypeInteractor
}

type MerchantsController interface {
	GetAll(c *fiber.Ctx) error
}

func NewMerchantsController(prototypeInteractor services.PrototypeInteractor) MerchantsController {
	return &merchantsController{
		prototypeInteractor,
	}
}
