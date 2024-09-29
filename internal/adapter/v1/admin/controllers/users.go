package controllers

import (
	"go-pudthai/internal/adapter/shared"
	"go-pudthai/internal/adapter/v1/admin/dtos"
	ia "go-pudthai/internal/usecase/interactor/admin"
	"go-pudthai/internal/validator"

	"github.com/gofiber/fiber/v2"
)

type usersController struct {
	userInteractor ia.UsersInteractor
}

func NewUsersController(userInteractor ia.UsersInteractor) UsersController {
	return &usersController{
		userInteractor,
	}
}

type UsersController interface {
	Create(c *fiber.Ctx) error
}

func (u usersController) Create(c *fiber.Ctx) error {
	req := dtos.CreateUser{}
	file, _ := c.FormFile("avatar")

	if errValidate := validator.Validate(c, &req); errValidate != nil {
		return errValidate
	}

	result, err := u.userInteractor.Create(c.Context(), req, file)
	if err != nil {
		return err
	}

	return shared.Success(c, result, nil, fiber.StatusCreated)
}
