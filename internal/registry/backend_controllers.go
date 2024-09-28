package registry

import (
	c "go-ibooking/internal/adapter/v1/backend/controllers"
)

func (r *registry) NewBackendAuthController() c.AuthController {
	return c.NewAuthController(r.NewSharedAuthInteractor())
}
