package router

import (
	controller "go-ibooking/internal/adapter/v1/controllers/admin"
	"go-ibooking/internal/model/technical"
)

func addAdminAuthRoute(c controller.AuthController) technical.Routes {
	return technical.Routes{
		technical.Route{
			Name: "Login", Method: "POST", Path: "login", Action: "Login", Subject: "Auth", HandlerFunc: c.Login,
		},
	}
}
