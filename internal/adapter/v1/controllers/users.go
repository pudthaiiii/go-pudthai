package controllers

import (
	"fmt"
	"go-ibooking/internal/model/technical"
	"go-ibooking/internal/usecase/interactor"
	"go-ibooking/internal/validator"

	"github.com/gofiber/fiber/v2"
)

type usersController struct {
	userInteractor interactor.UsersInteractor
}

func NewUsersController(userInteractor interactor.UsersInteractor) UsersController {
	return &usersController{
		userInteractor,
	}
}

type UsersController interface {
	Create(c *fiber.Ctx) error
}

func (u usersController) Create(c *fiber.Ctx) error {
	req := technical.CreateAdminUser{}
	file, _ := c.FormFile("avatar")

	fmt.Println(c.Locals("userId"))

	if errValidate := validator.Validate(c, &req); errValidate != nil {
		return errValidate
	}

	result, err := u.userInteractor.Create(c.Context(), req, file)
	if err != nil {
		return err
	}

	return success(c, result, nil)
}
