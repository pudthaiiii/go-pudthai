package router

import (
	cb "go-ibooking/internal/adapter/v1/backend/controllers"
	p "go-ibooking/internal/enum/permission"
	t "go-ibooking/internal/model/technical"
)

func addAuthRoute(c cb.AuthController) t.Routes {
	return t.Routes{
		// login
		t.Route{
			Name: "Login", Method: "POST", Path: "login", Action: string(p.NONE), Subject: string(p.AUTH), HandlerFunc: c.Login,
		},
		// refresh
		t.Route{
			Name: "Refresh", Method: "POST", Path: "refresh", Action: string(p.NONE), Subject: string(p.AUTH), HandlerFunc: c.Refresh,
		},
	}
}
