package router

import (
	controller "go-ibooking/internal/adapter/v1/admin/controllers"
	t "go-ibooking/internal/model/technical"
)

func addUsersRoute(c controller.UsersController) t.Routes {
	return t.Routes{
		t.Route{
			Name: "Create", Method: "POST", Path: "/users", Action: "", Subject: "", HandlerFunc: c.Create,
		},
	}
}
