package controllers

import (
	"go-ibooking/internal/enum"
	"go-ibooking/internal/model/dtos"
	"go-ibooking/internal/usecase/interactor"
	"go-ibooking/internal/validator"

	"github.com/gofiber/fiber/v2"
)

type authController struct {
	authInteractor interactor.AuthInteractor
}

func NewAuthController(authInteractor interactor.AuthInteractor) AuthController {
	return &authController{
		authInteractor,
	}
}

type AuthController interface {
	Login(c *fiber.Ctx) error
	LoginAdmin(c *fiber.Ctx) error
	LoginBackend(c *fiber.Ctx) error
}

func (s authController) Login(c *fiber.Ctx) error {
	req := dtos.Login{}
	if errValidate := validator.Validate(c, &req); errValidate != nil {
		return errValidate
	}

	result, err := s.authInteractor.Login(c.Context(), req, string(enum.USER))
	if err != nil {
		return err
	}

	return success(c, result, nil)
}

func (s authController) LoginAdmin(c *fiber.Ctx) error {
	req := dtos.Login{}
	if errValidate := validator.Validate(c, &req); errValidate != nil {
		return errValidate
	}

	result, err := s.authInteractor.Login(c.Context(), req, string(enum.ADMIN))
	if err != nil {
		return err
	}

	return success(c, result, nil)
}

func (s authController) LoginBackend(c *fiber.Ctx) error {
	req := dtos.Login{}
	if errValidate := validator.Validate(c, &req); errValidate != nil {
		return errValidate
	}

	result, err := s.authInteractor.Login(c.Context(), req, string(enum.MERCHANT))
	if err != nil {
		return err
	}

	return success(c, result, nil)
}
