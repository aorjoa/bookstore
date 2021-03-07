package main

import (
	"github.com/Aorjoa/bookstore/config"
	"github.com/Aorjoa/bookstore/router"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
)

func main() {
	db, err := config.DB()
	if err != nil {
		log.Panic().Err(err).Msg("cannot connect to database")
	}
	app := router.Init(db)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt)
	go func() {
		_ = <-sc
		log.Info().Msg("gracefully shutting down")
		_ = app.Shutdown()
	}()

	log.Info().Msg("server start listening")
	err = app.Listen(":3000")
	if err != nil {
		log.Panic().Err(err).Msg("server stop listening")
	}
	log.Info().Msg("cleanup task")
}
