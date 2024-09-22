package interactor

// import (
// 	"context"
// 	"fmt"
// 	"go-ibooking/internal/entities"
// 	"go-ibooking/internal/throw"
// 	"go-ibooking/internal/utils"
// 	"os"

// 	"golang.org/x/crypto/bcrypt"
// 	"gorm.io/gorm"
// )

// type authService struct {
// 	usersRepo    *gorm.DB
// 	usersService UsersService
// }

// type AuthService interface {
// 	Login(ctx context.Context, dto dtoReq.Login) (dtoRes.Create, error)
// }

// func NewAuthService(usersRepo *gorm.DB, usersService UsersService) AuthService {
// 	return &authService{
// 		usersRepo,
// 		usersService,
// 	}
// }

// // Create new user
// func (s *authService) Login(ctx context.Context, dto dtoReq.Login) (dtoRes.Create, error) {
// 	user, userErr := s.usersService.FindUserByEmail(ctx, dto.Email)
// 	if userErr != nil {
// 		return dtoRes.Create{}, userErr
// 	}

// 	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password)); err != nil {
// 		return dtoRes.Create{}, userErr
// 	}

// 	s.generateJwt(user)

// 	return dtoRes.Create{}, nil
// }

// func (s *authService) generateJwt(user entities.User) error {
// 	var (
// 		accessToken  string
// 		refreshToken string
// 		err          error
// 	)

// 	jwtAdminSecret := os.Getenv("JWT_ADMIN_SECRET")
// 	accessExpiresIn := os.Getenv("JWT_ACCESS_TOKEN_EXPIRES_IN_HOUR")
// 	refreshExpiresIn := os.Getenv("JWT_REFRESH_TOKEN_EXPIRES_IN_HOUR")

// 	accessToken, err = utils.JwtSign(map[string]interface{}{
// 		"userId": user.ID,
// 		"email":  user.Email,
// 	}, accessExpiresIn, jwtAdminSecret)
// 	if err != nil {
// 		return throw.Error(100001, err)
// 	}

// 	refreshToken, err = utils.JwtSign(map[string]interface{}{
// 		"userId": user.ID,
// 		"email":  user.Email,
// 	}, refreshExpiresIn, jwtAdminSecret)
// 	if err != nil {
// 		return throw.Error(100001, err)
// 	}

// 	fmt.Println(accessToken)
// 	fmt.Println(refreshToken)

// 	return nil
// }
