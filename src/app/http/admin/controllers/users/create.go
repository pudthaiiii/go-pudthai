package controllers

import (
	dtos "go-ibooking/src/app/http/admin/dtos/request/users"
	"go-ibooking/src/app/validator"

	"github.com/gofiber/fiber/v2"

	ApiResource "go-ibooking/src/app/resources"
)

func (s usersController) Create(c *fiber.Ctx) error {
	req := dtos.CreateRequest{}
	file, _ := c.FormFile("avatar")

	if errValidate := validator.Validate(c, &req); errValidate != nil {
		return errValidate
	}

	result, err := s.usersService.Create(c.Context(), req, file)
	if err != nil {
		return err
	}

	return ApiResource.SuccessResponse(c, result, nil)

}
