package router

import (
	controller "go-ibooking/internal/adapter/v1/controllers"
	"go-ibooking/internal/model/technical"
)

func addAuthRoute(c controller.AuthController) technical.Routes {
	return technical.Routes{
		// Login
		technical.Route{
			Name: "Login", Method: "POST", Path: "frontend/login", Action: "Login", Subject: "Auth", HandlerFunc: c.Login,
		},

		// Login Backend
		technical.Route{
			Name: "LoginBackend", Method: "POST", Path: "backend/login", Action: "Login", Subject: "Auth", HandlerFunc: c.LoginBackend,
		},

		// Login Admin
		technical.Route{
			Name: "LoginAdmin", Method: "POST", Path: "admin/login", Action: "Login", Subject: "Auth", HandlerFunc: c.LoginAdmin,
		},
	}
}
