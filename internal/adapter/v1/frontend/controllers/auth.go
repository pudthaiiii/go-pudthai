package controllers

import (
	"go-ibooking/internal/adapter/shared"
	"go-ibooking/internal/adapter/shared/dtos"
	t "go-ibooking/internal/model/technical"
	is "go-ibooking/internal/usecase/interactor/shared"
	"go-ibooking/internal/validator"

	"github.com/gofiber/fiber/v2"
)

type authController struct {
	authInteractor is.SharedAuthInteractor
}

func NewAuthController(authInteractor is.SharedAuthInteractor) AuthController {
	return &authController{
		authInteractor,
	}
}

type AuthController interface {
	Login(c *fiber.Ctx) error
	Refresh(c *fiber.Ctx) error
}

func (s authController) Login(c *fiber.Ctx) error {
	dto := dtos.Login{}
	if errValidate := validator.Validate(c, &dto); errValidate != nil {
		return errValidate
	}

	result, err := s.authInteractor.Login(c.Context(), dto, string(t.USER))
	if err != nil {
		return err
	}

	return shared.Success(c, result, nil)
}

func (s authController) Refresh(c *fiber.Ctx) error {
	dto := dtos.Refresh{}
	if errValidate := validator.Validate(c, &dto); errValidate != nil {
		return errValidate
	}

	result, err := s.authInteractor.Refresh(c.Context(), dto, string(t.USER))
	if err != nil {
		return err
	}

	return shared.Success(c, result, nil)
}
