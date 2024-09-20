package registry

import (
	c "go-ibooking/internal/adapter/v1/controllers"
)

// UsersController
func (r *registry) RegisterUsersController() c.UsersController {
	return c.NewUsersController(r.RegisterUsersInteractor())
}
