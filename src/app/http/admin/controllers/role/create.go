package controllers

import (
	dtos "go-ibooking/src/app/http/admin/dtos/request/roles"
	ApiResource "go-ibooking/src/app/http/resources"
	"go-ibooking/src/app/validator"

	"github.com/gofiber/fiber/v2"
)

func (s roleController) Create(c *fiber.Ctx) error {
	req := dtos.RoleCreateRequest{}

	if errValidate := validator.Validate(c, &req); errValidate != nil {
		return errValidate
	}

	result, err := s.roleService.Create(c.Context(), req)
	if err != nil {
		return err
	}

	return ApiResource.SuccessResponse(c, result, nil)

}
