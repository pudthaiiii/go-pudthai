package registry

import (
	c "go-pudthai/internal/adapter/v1/frontend/controllers"
)

func (r *registry) NewFrontendAuthController() c.AuthController {
	return c.NewAuthController(r.NewSharedAuthInteractor())
}
