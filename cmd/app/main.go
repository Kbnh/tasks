package main

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/Kbnh/tasks/config"
	"github.com/Kbnh/tasks/internal/app"
)

func main() {
	c, err := config.New()
	if err != nil {
		log.Fatal().Err(err).Msg("config.New")
	}

	ctx := context.Background()

	err = app.Run(ctx, c)
	if err != nil {
		log.Fatal().Err(err).Msg("app.Run")
	}
}
