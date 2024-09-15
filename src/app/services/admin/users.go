package services

import (
	"context"
	"fmt"
	"mime/multipart"

	pkg "go-ibooking/src/pkg"

	"github.com/google/uuid"
	"gorm.io/gorm"

	dtoReq "go-ibooking/src/app/http/admin/dtos/request/users"
	dtoRes "go-ibooking/src/app/http/admin/dtos/response/users"
)

type UsersService interface {
	Create(ctx context.Context, dto dtoReq.CreateRequest, avatar *multipart.FileHeader) (dtoRes.CreateResponse, error)
}

type usersService struct {
	usersRepo *gorm.DB
	s3        *pkg.S3Datastore
}

func NewUsersService(usersRepo *gorm.DB, s3 *pkg.S3Datastore) UsersService {
	return &usersService{
		usersRepo,
		s3,
	}
}

// Create godoc
// @Summary Create a new user
// @Description Creates a new user and optionally uploads an avatar
// @Tags users
// @Accept  multipart/form-data
// @Produce  json
// @Param dto body dtoReq.CreateRequest true "User data"
// @Param avatar formData file false "User avatar"
// @Success 200 {object} dtoRes.CreateResponse
// @Router /users [post]
func (s *usersService) Create(ctx context.Context, dto dtoReq.CreateRequest, avatar *multipart.FileHeader) (dtoRes.CreateResponse, error) {
	response := dtoRes.CreateResponse{}
	fileName := ""

	// Upload avatar to S3
	if avatar != nil {
		avatarName := uuid.New()
		fileName = fmt.Sprintf("users/%s%s", avatarName.String(), ".jpg")

		_, err := s.s3.ValidateAndUpload(avatar, fileName)
		if err != nil {
			return response, err
		}
	}

	return response, nil
}
