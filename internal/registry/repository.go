package registry

import (
	"go-pudthai/internal/entities"
	"go-pudthai/internal/usecase/repository"

	"gorm.io/gorm"
)

func (r *registry) NewUsersRepository() repository.UsersRepository {
	return repository.NewUsersRepository(
		r.db.Model(&entities.User{}).Session(&gorm.Session{NewDB: true}),
	)
}

func (r *registry) NewOauthAccessTokenRepository() repository.OauthAccessTokenRepository {
	return repository.NewOauthAccessTokenRepository(
		r.db.Model(&entities.OauthAccessToken{}).Session(&gorm.Session{NewDB: true}),
	)
}

func (r *registry) NewOauthRefreshTokenRepository() repository.OauthRefreshTokenRepository {
	return repository.NewOauthRefreshTokenRepository(
		r.db.Model(&entities.OauthRefreshToken{}).Session(&gorm.Session{NewDB: true}),
	)
}
