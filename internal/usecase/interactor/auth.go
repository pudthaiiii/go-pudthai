package interactor

import (
	"context"
	"fmt"
	"go-ibooking/internal/config"
	"go-ibooking/internal/infrastructure/cache"
	"go-ibooking/internal/model/dtos"
	"go-ibooking/internal/throw"
	"go-ibooking/internal/usecase/repository"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type authInteractor struct {
	userRepo     repository.UsersRepository
	cacheManager *cache.CacheManager
	cfg          *config.Config
}

func NewAuthInteractor(usersRepo repository.UsersRepository, cacheManager *cache.CacheManager, cfg *config.Config) AuthInteractor {
	return &authInteractor{
		usersRepo,
		cacheManager,
		cfg,
	}
}

type AuthInteractor interface {
	Login(ctx context.Context, dto dtos.Login, userType string) (dtos.AuthJWT, error)
}

// Create new user
func (s *authInteractor) Login(ctx context.Context, dto dtos.Login, userType string) (dtos.AuthJWT, error) {
	user, userErr := s.userRepo.FindUserByEmail(ctx, dto.Email, userType)
	if userErr != nil {
		return dtos.AuthJWT{}, throw.UserCredentialMismatch()
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password)); err != nil {
		return dtos.AuthJWT{}, throw.UserCredentialMismatch()
	}

	payload := map[string]interface{}{
		"user_id": user.ID,
		"role":    "admin",
	}

	_, errJwt := s.generateJwt(payload, userType)
	if errJwt != nil {
		return dtos.AuthJWT{}, throw.UserCredentialMismatch()
	}

	return dtos.AuthJWT{}, nil
}

func (s *authInteractor) generateJwt(payload map[string]interface{}, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"iat": time.Now().Unix(),
	})

	claims := token.Claims.(jwt.MapClaims)

	for key, value := range payload {
		claims[key] = value
	}

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	fmt.Println(tokenString)

	return tokenString, nil
}

// func (s *authInteractor) generateJwt(user entities.User, userType string) error {
// 	var (
// 		accessToken string
// 		// refreshToken string
// 		err    error
// 		secret string
// 	)

// 	switch userType {
// 	case string(enum.USER):
// 		secret = s.cfg.Get("JWT")["JwtSecret"].(string)
// 	case string(enum.ADMIN):
// 		secret = s.cfg.Get("JWT")["JwtSecretAdmin"].(string)
// 	case string(enum.MERCHANT):
// 		secret = s.cfg.Get("JWT")["JwtSecretBackend"].(string)
// 	}

// 	accessExpiresIn := s.cfg.Get("JWT")["AccessTokenExpiresInHour"].(string)
// 	// refreshExpiresIn := s.cfg.Get("JWT")["JwtRefreshExpiresInHour"].(string)

// 	payload := map[string]interface{}{
// 		"userId": user.ID,
// 		"email":  user.Email,
// 		"role":   userType,
// 	}

// 	accessToken, err = utils.JwtSign(payload, accessExpiresIn, secret)
// 	if err != nil {
// 		return throw.UserCredentialMismatch()
// 	}

// 	// refreshToken, err = utils.JwtSign(map[string]interface{}{
// 	// 	"userId": user.ID,
// 	// 	"email":  user.Email,
// 	// }, refreshExpiresIn, secret)
// 	// if err != nil {
// 	// 	return throw.UserCredentialMismatch()
// 	// }

// 	fmt.Println(accessToken)
// 	// fmt.Println(refreshToken)

// 	return nil
// }
