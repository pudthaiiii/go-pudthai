package main

import (
	"go-pudthai/cmd/api"
)

func main() {
	application := api.NewApiApplication()
	application.DeferClose()

	application.Boot()
}
