package controllers

import (
	"go-ibooking/internal/adapter/shared"
	"go-ibooking/internal/model/dtos"
	ia "go-ibooking/internal/usecase/interactor/admin"
	"go-ibooking/internal/validator"

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

// Create godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags Users
// @Accept json
// @Produce json
// @Param avatar formData file false "User avatar"
// @Param email formData string true "User email"
// @Param password formData string true "User password"
// @Success 200 {object} dtos.ResponseUserID
// @Router /v1/users [post]
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

	return shared.Success(c, result, nil)
}
