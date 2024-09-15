package controllers

import (
	service "go-ibooking/src/app/services/admin"

	"github.com/gofiber/fiber/v2"
)

type RoleController interface {
	Paginate(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
}

type roleController struct {
	roleService service.RolesService
}

func NewRoleController(roleService service.RolesService) RoleController {
	return &roleController{
		roleService,
	}
}
