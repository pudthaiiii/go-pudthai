package registry

import (
	c "go-ibooking/internal/adapter/v1/controllers"
)

// UsersController
func (r *registry) NewUsersController() c.UsersController {
	return c.NewUsersController(r.NewUsersInteractor())
}

// AuthController
func (r *registry) NewAuthController() c.AuthController {
	return c.NewAuthController(r.NewSharedAuthInteractor())
}
