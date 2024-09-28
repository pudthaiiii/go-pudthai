package registry

import (
	i "go-ibooking/internal/usecase/interactor/admin"
)

func (r *registry) NewAdminUsersInteractor() i.UsersInteractor {
	return i.NewUsersInteractor(
		r.NewUsersRepository(),
		r.s3,
		r.listener,
	)
}
