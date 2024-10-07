package interactor

import (
	"context"
	"errors"
	"fmt"
	"go-pudthai/internal/adapter/v1/admin/dtos"
	"go-pudthai/internal/entities"
	"go-pudthai/internal/events"
	"go-pudthai/internal/infrastructure/datastore"
	"go-pudthai/internal/model/business"
	t "go-pudthai/internal/model/technical"
	"go-pudthai/internal/throw"
	"go-pudthai/internal/usecase/repository"
	"mime/multipart"
	"strings"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type usersInteractor struct {
	userRepo repository.UsersRepository
	s3       *datastore.S3Datastore
	listener events.EventListener
}

func NewUsersInteractor(userRepo repository.UsersRepository, s3 *datastore.S3Datastore, listener events.EventListener) UsersInteractor {
	return &usersInteractor{
		userRepo,
		s3,
		listener,
	}
}

type UsersInteractor interface {
	Create(ctx context.Context, dto dtos.CreateUser, avatar *multipart.FileHeader) (dtos.ResponseUserID, error)
}

func (u *usersInteractor) Create(ctx context.Context, dto dtos.CreateUser, file *multipart.FileHeader) (response dtos.ResponseUserID, err error) {
	var (
		createUser entities.User
	)

	fileName, err := u.handleFileUpload(ctx, file)
	if err != nil {
		return response, err
	}

	userInfo := ctx.Value(t.UserInfo).(business.UserInfo)
	if userInfo.Type != string(t.ADMIN) {
		dto.MerchantID = userInfo.MerchantID
	}

	if existingUser, err := u.userRepo.FindUserByEmail(ctx, dto.Email, dto.Type); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return response, throw.UserCreate(err)
	} else if existingUser.ID != 0 {
		return response, throw.UserExists()
	}

	hashedPassword, err := hashPassword(dto.Password)
	if err != nil {
		return response, throw.UserCreate(err)
	}

	copier.Copy(&createUser, &dto)

	createUser.Email = strings.ToLower(dto.Email)
	createUser.Password = string(hashedPassword)
	createUser.ProfileImage = fileName

	savedUser, err := u.userRepo.CreateAdminUser(ctx, createUser)
	if err != nil {
		return response, throw.UserCreate(err)
	}

	copier.Copy(&response, &savedUser)

	u.emitUserCreatedEvent(dto.Type, savedUser)

	return response, nil
}

func (u *usersInteractor) handleFileUpload(ctx context.Context, file *multipart.FileHeader) (string, error) {
	if file == nil {
		return "", nil
	}
	fileName := fmt.Sprintf("users/%s.jpg", uuid.New().String())
	if _, err := u.s3.ValidateAndUpload(ctx, file, fileName); err != nil {
		return "", throw.UploadError(err)
	}
	return fileName, nil
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func (u *usersInteractor) emitUserCreatedEvent(userType string, user interface{}) {
	switch t.UserTypeEnum(strings.ToUpper(userType)) {
	case t.ADMIN:
		events.Emit(u.listener, "admin.created", user)
	case t.MERCHANT:
		events.Emit(u.listener, "merchant.created", user)
	default:
		events.Emit(u.listener, "user.created", user)
	}
}
