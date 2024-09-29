package router

import (
	controller "go-pudthai/internal/adapter/v1/admin/controllers"
	t "go-pudthai/internal/model/technical"
)

func addUsersRoute(c controller.UsersController) t.Routes {
	return t.Routes{
		t.Route{
			Name: "Create", Method: "POST", Path: "/users", Action: string(t.MANAGER), Subject: string(t.USER), HandlerFunc: c.Create,
		},
	}
}
