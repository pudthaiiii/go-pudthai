package controllers

import (
	dtos "go-ibooking/src/app/http/admin/dtos/request/roles"
	ApiResource "go-ibooking/src/app/resources"
	"go-ibooking/src/app/validator"

	"github.com/gofiber/fiber/v2"
)

func (s roleController) Create(c *fiber.Ctx) error {
	var req dtos.RoleCreateRequest

	if errValidate := validator.Validate(c, &req); errValidate != nil {
		return errValidate
	}

	response, err := s.roleService.Create(c.Context(), req)
	if err != nil {
		return err
	}

	return ApiResource.SuccessResponse(c, response, nil, nil)

}
