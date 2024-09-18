package controllers

import (
	dtos "go-ibooking/src/app/http/admin/dtos/request"
	ApiResource "go-ibooking/src/app/http/resources"
	"go-ibooking/src/app/validator"

	"github.com/gofiber/fiber/v2"
)

func (s roleController) Paginate(c *fiber.Ctx) error {
	params := dtos.PaginateRequest{}
	params.SetDefaults()

	if errValidate := validator.Validate(c, &params); errValidate != nil {
		return errValidate
	}

	result, paginate, err := s.roleService.Paginate(c.Context(), params)
	if err != nil {
		return err
	}

	return ApiResource.SuccessResponse(c, result, paginate)
}
