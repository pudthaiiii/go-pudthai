package registry

import (
	i "go-ibooking/internal/usecase/interactor"
)

func (r *registry) NewUsersInteractor() i.UsersInteractor {
	return i.NewUsersInteractor(
		r.NewUsersRepository(),
		r.s3,
		r.listener,
	)
}

func (r *registry) NewAuthInteractor() i.AuthInteractor {
	return i.NewAuthInteractor(
		r.NewUsersRepository(),
		r.cacheManager,
		r.cfg,
	)
}
