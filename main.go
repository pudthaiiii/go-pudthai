package main

import "go-ibooking/cmd/api"

func main() {
	application := api.NewApiApplication()
	application.DeferClose()

	application.Boot()
}
