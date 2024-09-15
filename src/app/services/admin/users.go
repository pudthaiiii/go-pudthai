package services

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"

	pkg "go-ibooking/src/pkg"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	throw "go-ibooking/src/app/exception"
	dtoReq "go-ibooking/src/app/http/admin/dtos/request/users"
	dtoRes "go-ibooking/src/app/http/admin/dtos/response/users"
	"go-ibooking/src/app/model"
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

// Create new user
func (s *usersService) Create(ctx context.Context, dto dtoReq.CreateRequest, avatar *multipart.FileHeader) (dtoRes.CreateResponse, error) {
	fileName := ""
	response := dtoRes.CreateResponse{}

	// check existing user
	existingErr := s.existingUserByEmail(dto.Email)
	if existingErr != nil {
		return response, existingErr
	}

	// upload avatar
	if avatar != nil {
		avatarName := uuid.New()
		fileName = fmt.Sprintf("users/%s%s", avatarName.String(), ".jpg")

		_, err := s.s3.ValidateAndUpload(avatar, fileName)
		if err != nil {
			return response, err
		}
	}

	// hash password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)

	// create user
	user := model.User{
		Email:        dto.Email,
		Password:     string(hashedPassword),
		RoleID:       dto.RoleID,
		ProfileImage: fileName,
		IsActive:     dto.IsActive,
		IsAllBU:      dto.IsAllBU,
		FullName:     dto.FullName,
		Mobile:       dto.Mobile,
		Company:      dto.Company,
		Type:         dto.Type,
		MerchantID:   1,
	}

	// save user
	createUser := s.usersRepo.Create(&user)
	if createUser.Error != nil {
		return response, throw.Error(910201, createUser.Error)
	}

	// response
	response = dtoRes.CreateResponse{
		ID: user.ID,
	}

	return response, nil
}

// Find user by email
func (s *usersService) existingUserByEmail(email string) error {
	user := model.User{
		Email: email,
	}

	userExists := s.usersRepo.Where(&user).First(&user)
	if userExists.Error == nil {
		return throw.Error(910202, nil)
	}

	if !errors.Is(userExists.Error, gorm.ErrRecordNotFound) {
		return throw.Error(910202, userExists.Error)
	}

	return nil
}
