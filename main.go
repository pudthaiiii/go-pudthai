package main

import (
	"fmt"
	"log"

	"github.com/gocraft/work"
	"github.com/gofiber/fiber/v2"
	"github.com/gomodule/redigo/redis"
	"github.com/pudthaiiii/go-ibooking/src/cmd"
	resource "github.com/pudthaiiii/go-ibooking/src/resource"
)

var RedisPool *redis.Pool
var Enqueuer *work.Enqueuer

func main() {
	RedisPool = &redis.Pool{
		MaxActive: 5,
		MaxIdle:   5,
		Wait:      true,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", ":6379")
		},
	}

	Enqueuer = work.NewEnqueuer("go-ibooking", RedisPool)

	app := initFiberRouter()

	application := cmd.NewApplication(app)

	application.Boot()

	log.Fatal(app.Listen(":3000"))
}

func initFiberRouter() *fiber.App {
	cmd.InitializeEnv()

	app := fiber.New(fiber.Config{
		ErrorHandler: resource.ErrorHandler,
	})

	_, err := Enqueuer.Enqueue("test_job", work.Q{"name": "send by init"})
	if err != nil {
		fmt.Println(err)
	}

	// Import job queue routes
	initJobQueueRoutes(app)

	return app
}
