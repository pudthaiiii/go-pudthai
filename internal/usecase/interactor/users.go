package interactor

import (
	"context"
	"fmt"
	"mime/multipart"

	"go-ibooking/internal/infrastructure/datastore"
	"go-ibooking/internal/model/dtos"
	"go-ibooking/internal/usecase/repository"

	"github.com/google/uuid"

	throw "go-ibooking/internal/exception"
)

type usersInteractor struct {
	userRepo repository.UsersRepository
	s3       *datastore.S3Datastore
}

type UsersInteractor interface {
	Create(ctx context.Context, dto dtos.CreateUser, avatar *multipart.FileHeader) (dtos.ShowUser, error)
	// FindUserByEmail(ctx context.Context, email string) (entities.User, error)
}

func NewUsersInteractor(userRepo repository.UsersRepository, s3 *datastore.S3Datastore) UsersInteractor {
	return &usersInteractor{
		userRepo,
		s3,
	}
}

// Create new user
func (u *usersInteractor) Create(ctx context.Context, dto dtos.CreateUser, file *multipart.FileHeader) (dtos.ShowUser, error) {
	var (
		err        error
		fileName   string
		createUser dtos.ShowUser
	)

	// check existing user
	// existingErr := s.existingUserByEmail(dto.Email)
	// if existingErr != nil {
	// 	return response, existingErr
	// }

	// upload avatar
	if file != nil {
		avatarName := uuid.New()
		fileName = fmt.Sprintf("users/%s%s", avatarName.String(), ".jpg")

		_, err := u.s3.ValidateAndUpload(ctx, file, fileName)
		if err != nil {
			return createUser, throw.Error(910201, err)
		}
	}

	// save user
	createUser, err = u.userRepo.CreateAdminUser(ctx, dto, fileName)
	if err != nil {
		return createUser, throw.Error(910201, err)
	}

	return createUser, nil
}

// // check exist user by email
// func (s *usersInteractor) existingUserByEmail(email string) error {
// 	user := entities.User{
// 		Email: email,
// 	}

// 	userExists := s.usersRepo.Where(&user).First(&user)
// 	if userExists.Error == nil {
// 		return throw.Error(910202, nil)
// 	}

// 	if !errors.Is(userExists.Error, gorm.ErrRecordNotFound) {
// 		return throw.Error(910202, userExists.Error)
// 	}

// 	return nil
// }

// func (s *usersInteractor) FindUserByEmail(ctx context.Context, email string) (entities.User, error) {
// 	user := entities.User{}

// 	userQuery := s.usersRepo.Where("email = ?", email).First(&user)
// 	if userQuery.Error == gorm.ErrRecordNotFound {
// 		return user, throw.Error(910204, userQuery.Error)
// 	}

// 	return user, nil
// }
