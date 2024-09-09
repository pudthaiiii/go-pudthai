package types

import "github.com/gofiber/fiber/v2"

type Route struct {
	Name         string
	Method       string
	Pattern      string
	Operation    string
	Resource     string
	HandlerFunc  fiber.Handler
	IsPathPrefix bool
}

type Routes []Route
