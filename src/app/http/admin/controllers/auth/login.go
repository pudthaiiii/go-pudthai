package controllers

import (
	dtos "go-ibooking/src/app/http/admin/dtos/request/auth"
	ApiResource "go-ibooking/src/app/resources"
	"go-ibooking/src/app/validator"

	"github.com/gofiber/fiber/v2"
)

func (s authController) Login(c *fiber.Ctx) error {
	req := dtos.Login{}

	if errValidate := validator.Validate(c, &req); errValidate != nil {
		return errValidate
	}

	result, err := s.authService.Login(c.Context(), req)
	if err != nil {
		return err
	}

	// tokenG := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"userId": 1,
	// 	"uuid":   "5a22cda8-141e-4237-b1a5-3cd7e4ad664d",
	// 	"exp":    time.Now().Add(time.Minute * 5).Unix(),
	// 	"iat":    time.Now().Unix(),
	// })

	// fmt.Println(tokenG.SignedString([]byte(os.Getenv("JWT_ADMIN_SECRET"))))

	return ApiResource.SuccessResponse(c, result, nil)
}
