package route

import (
	controller "go-ibooking/src/app/http/admin/controllers/auth"
	"go-ibooking/src/types"
)

func addAuthRoute(c controller.AuthController) types.Routes {
	return types.Routes{
		types.Route{
			Name:        "Login",
			Method:      "POST",
			Path:        "/login",
			Action:      "Login",
			Subject:     "Auth",
			HandlerFunc: c.Login,
		},
	}
}
