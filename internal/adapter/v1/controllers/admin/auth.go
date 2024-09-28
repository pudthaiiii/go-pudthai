package controllers

import (
	"go-ibooking/internal/adapter/v1/controllers"
	"go-ibooking/internal/enum"
	"go-ibooking/internal/model/dtos"
	i "go-ibooking/internal/usecase/interactor/shared"
	"go-ibooking/internal/validator"

	"github.com/gofiber/fiber/v2"
)

type authController struct {
	authInteractor i.SharedAuthInteractor
}

func NewAuthController(authInteractor i.SharedAuthInteractor) AuthController {
	return &authController{
		authInteractor,
	}
}

type AuthController interface {
	Login(c *fiber.Ctx) error
}

func (s authController) Login(c *fiber.Ctx) error {
	req := dtos.Login{}
	if errValidate := validator.Validate(c, &req); errValidate != nil {
		return errValidate
	}

	result, err := s.authInteractor.Login(c.Context(), req, string(enum.ADMIN))
	if err != nil {
		return err
	}

	return controllers.Success(c, result, nil)
}
