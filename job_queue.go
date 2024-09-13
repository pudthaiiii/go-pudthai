package main

import (
	"log"
	"time"

	"github.com/gocraft/work"
	"github.com/gofiber/fiber/v2"
)

func initJobQueueRoutes(app *fiber.App) {
	pool := work.NewWorkerPool(Context{}, 10, "go-ibooking", RedisPool)

	pool.Job("test_job", (*Context).TestJob)

	go pool.Start()
	defer pool.Stop()

	app.Post("/enqueue", func(c *fiber.Ctx) error {
		_, err := Enqueuer.Enqueue("test_job", work.Q{"name": "send by post mans"})
		if err != nil {
			return err
		}

		return c.SendString("Job type 1 enqueued")
	})

}

type Context struct{}

func (c *Context) TestJob(job *work.Job) error {
	name := job.ArgString("name")

	if err := job.ArgError(); err != nil {
		return err
	}

	log.Printf("Processing job type 1 with name: %s", name)
	time.Sleep(2 * time.Second)
	return nil
}
