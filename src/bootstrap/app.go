package bootstrap

func Boot() {
	initializeEnv()

	router := routerInit()

	start(router)
}
