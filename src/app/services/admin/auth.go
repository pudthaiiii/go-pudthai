package services

import (
	"context"
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	throw "go-ibooking/src/app/exception"
	dtoReq "go-ibooking/src/app/http/admin/dtos/request/auth"
	dtoRes "go-ibooking/src/app/http/admin/dtos/response/users"
	"go-ibooking/src/app/model"
	"go-ibooking/src/utils"
)

type authService struct {
	usersRepo    *gorm.DB
	usersService UsersService
}

type AuthService interface {
	Login(ctx context.Context, dto dtoReq.Login) (dtoRes.Create, error)
}

func NewAuthService(usersRepo *gorm.DB, usersService UsersService) AuthService {
	return &authService{
		usersRepo,
		usersService,
	}
}

// Create new user
func (s *authService) Login(ctx context.Context, dto dtoReq.Login) (dtoRes.Create, error) {
	user, userErr := s.usersService.FindUserByEmail(ctx, dto.Email)
	if userErr != nil {
		return dtoRes.Create{}, throw.Error(100001, nil)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password)); err != nil {
		return dtoRes.Create{}, throw.Error(100001, nil)
	}

	s.generateJwt(user)

	return dtoRes.Create{}, nil
}

func (s *authService) generateJwt(user model.User) error {

	jwtAdminSecret := os.Getenv("JWT_ADMIN_SECRET")
	accessExpiresIn := os.Getenv("JWT_ACCESS_TOKEN_EXPIRES_IN_HOUR")
	refreshExpiresIn := os.Getenv("JWT_REFRESH_TOKEN_EXPIRES_IN_HOUR")

	accessToken, err := utils.JwtSign(map[string]interface{}{
		"userId": user.ID,
		"email":  user.Email,
	}, accessExpiresIn, jwtAdminSecret)

	refreshToken, err := utils.JwtSign(map[string]interface{}{
		"userId": user.ID,
		"email":  user.Email,
	}, refreshExpiresIn, jwtAdminSecret)

	if err != nil {
		return throw.Error(100001, err)
	}

	fmt.Println(accessToken)
	fmt.Println(refreshToken)

	return nil
}
