package interactor

import (
	"context"
	"errors"
	"fmt"
	"go-ibooking/internal/events"
	"go-ibooking/internal/infrastructure/datastore"
	"go-ibooking/internal/model/dtos"
	"go-ibooking/internal/throw"
	"go-ibooking/internal/usecase/repository"
	"mime/multipart"

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

type UsersInteractor interface {
	Create(ctx context.Context, dto dtos.CreateUser, avatar *multipart.FileHeader) (dtos.ResponseUserID, error)
}

func NewUsersInteractor(userRepo repository.UsersRepository, s3 *datastore.S3Datastore, listener events.EventListener) UsersInteractor {
	return &usersInteractor{
		userRepo,
		s3,
		listener,
	}
}

func (u *usersInteractor) Create(ctx context.Context, dto dtos.CreateUser, file *multipart.FileHeader) (dtos.ResponseUserID, error) {
	var (
		createUser dtos.ResponseUserID
		fileName   string
	)
	events.Emit(u.listener, "create_user_email", nil)

	return createUser, nil
	existingUser, err := u.userRepo.FindUserByEmail(ctx, dto.Email, "")
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return createUser, throw.UserCreate(err)
		}
	} else if existingUser.ID != 0 {
		return createUser, throw.UserExists()
	}

	if file != nil {
		fileName = fmt.Sprintf("users/%s.jpg", uuid.New().String())
		if _, err := u.s3.ValidateAndUpload(ctx, file, fileName); err != nil {
			return createUser, throw.UploadError(err)
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return createUser, throw.UserCreate(err)
	}

	user, err := u.userRepo.CreateAdminUser(ctx, dto, fileName, string(hashedPassword))
	if err != nil {
		return createUser, throw.UserCreate(err)
	}

	events.Emit(u.listener, "create_user_email", user)

	copier.Copy(&createUser, &user)

	return createUser, nil
}
