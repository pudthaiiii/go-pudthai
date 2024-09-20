package main

import (
	"go-ibooking/cmd/app"
)

type Abc struct {
	Port string
}

func main() {
	application := app.NewApplication()

	application.Boot()

	application.Listen()
}
