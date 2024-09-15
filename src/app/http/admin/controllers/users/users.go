package controllers

import (
	service "go-ibooking/src/app/services/admin"

	"github.com/gofiber/fiber/v2"
)

type UsersController interface {
	Create(c *fiber.Ctx) error
}

type usersController struct {
	usersService service.UsersService
}

func NewUsersController(usersService service.UsersService) UsersController {
	return &usersController{
		usersService,
	}
}
