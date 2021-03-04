package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
)

func main() {
	app := fiber.New()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt)
	go func() {
		_ = <-sc
		log.Info().Msg("gracefully shutting down")
		_ = app.Shutdown()
	}()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, world :)")
	})

	log.Info().Msg("server start listening")
	err := app.Listen(":3000")
	if err != nil {
		log.Panic().Err(err).Msg("server stop listening")
	}
	log.Info().Msg("cleanup task")
}
