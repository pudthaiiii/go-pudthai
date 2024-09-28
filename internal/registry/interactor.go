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
