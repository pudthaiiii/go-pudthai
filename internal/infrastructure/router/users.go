package router

import (
	controller "go-ibooking/internal/adapter/v1/controllers"
	"go-ibooking/internal/model/technical"
)

func addUsersRoute(c controller.UsersController) technical.Routes {
	return technical.Routes{
		technical.Route{
			Name:        "Create",
			Method:      "POST",
			Path:        "/users",
			Action:      "",
			Subject:     "",
			HandlerFunc: c.Create,
		},
	}
}
