package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"os/signal"
)

func main() {
	app := fiber.New()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt)
	go func() {
		_ = <-sc
		log.Println("gracefully shutting down")
		_ = app.Shutdown()
	}()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, world :)")
	})

	err := app.Listen(":3000")
	if err != nil {
		log.Panic(err)
	}
	log.Println("cleanup task")
}
