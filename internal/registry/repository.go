package registry

import (
	"go-ibooking/internal/entities"
	"go-ibooking/internal/usecase/repository"

	"gorm.io/gorm"
)

func (r *registry) NewUsersRepository() repository.UsersRepository {
	return repository.NewUsersRepository(
		r.db.Model(&entities.User{}).Session(&gorm.Session{NewDB: true}),
	)
}
