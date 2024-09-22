package main

import (
	"go-ibooking/cmd/api"
	"go-ibooking/internal/events"
)

func main() {
	listener := events.NewEventListener()

	// เริ่ม listener ใน goroutine
	go listener.Listen()

	events.Emit(listener, "user_signup", map[string]string{"username": "john_doe"})

	application := api.NewApiApplication()
	application.DeferClose()

	application.Boot()
}
