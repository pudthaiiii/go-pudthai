package technical

import "github.com/gofiber/fiber/v2"

type Route struct {
	Name         string
	Method       string
	Path         string
	Action       string
	Subject      string
	HandlerFunc  fiber.Handler
	IsPathPrefix bool
}

type Routes []Route
