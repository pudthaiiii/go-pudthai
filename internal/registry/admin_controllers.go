package registry

import (
	c "go-pudthai/internal/adapter/v1/admin/controllers"
)

// AuthController
func (r *registry) NewAdminAuthController() c.AuthController {
	return c.NewAuthController(r.NewSharedAuthInteractor())
}

func (r *registry) NewAdminUsersController() c.UsersController {
	return c.NewUsersController(r.NewAdminUsersInteractor())
}
