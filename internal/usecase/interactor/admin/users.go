package interactor

import (
	"context"
	"errors"
	"fmt"
	"go-pudthai/internal/adapter/v1/admin/dtos"
	"go-pudthai/internal/events"
	"go-pudthai/internal/infrastructure/datastore"
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

func (u *usersInteractor) Create(ctx context.Context, dto dtos.CreateUser, file *multipart.FileHeader) (dtos.ResponseUserID, error) {

	fmt.Println(ctx.Value(t.Merchant),
		ctx.Value(t.MerchantID),
		ctx.Value(t.Member))
	var createUser dtos.ResponseUserID

	userType, merchantID := resolveUserTypeAndMerchantID(dto.Type)

	if existingUser, err := u.userRepo.FindUserByEmail(ctx, dto.Email, userType); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return createUser, throw.UserCreate(err)
	} else if existingUser.ID != 0 {
		return createUser, throw.UserExists()
	}

	fileName, err := u.handleFileUpload(ctx, file)
	if err != nil {
		return createUser, err
	}

	hashedPassword, err := hashPassword(dto.Password)
	if err != nil {
		return createUser, throw.UserCreate(err)
	}

	user, err := u.userRepo.CreateAdminUser(ctx, dto, fileName, hashedPassword, merchantID, userType)
	if err != nil {
		return createUser, throw.UserCreate(err)
	}

	copier.Copy(&createUser, &user)

	u.emitUserCreatedEvent(dto.Type, user)

	return createUser, nil
}

func resolveUserTypeAndMerchantID(userType string) (string, uint) {
	switch t.UserTypeEnum(strings.ToUpper(userType)) {
	case t.ADMIN:
		return string(t.ADMIN), 0
	case t.MERCHANT:
		return string(t.MERCHANT), 99
	default:
		return "User", 99
	}
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
