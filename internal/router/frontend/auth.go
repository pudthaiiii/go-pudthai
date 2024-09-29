package router

import (
	ac "go-pudthai/internal/adapter/v1/frontend/controllers"
	t "go-pudthai/internal/model/technical"
)

func addAuthRoute(c ac.AuthController) t.Routes {
	return t.Routes{
		// login
		t.Route{
			Name: "Login", Method: "POST", Path: "login", Action: string(t.NONE), Subject: string(t.AUTH), HandlerFunc: c.Login,
		},
		// refresh
		t.Route{
			Name: "Refresh", Method: "POST", Path: "refresh", Action: string(t.NONE), Subject: string(t.AUTH), HandlerFunc: c.Refresh,
		},
	}
}
