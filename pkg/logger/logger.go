package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Config struct {
	AppName    string `envconfig:"APP_NAME" required:"true"`
	AppVersion string `envconfig:"APP_VERSION" required:"true"`
	Level      string `default:"error" envconfig:"LOGGER_LEVEL"`
	Pretty     bool   `default:"false" envconfig:"LOGGER_PRETTY"`
}

func Init(c Config) {
	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	if level, err := zerolog.ParseLevel(c.Level); err == nil && level != zerolog.NoLevel {
		zerolog.SetGlobalLevel(level)
	}

	log.Logger = log.With().
		Str("app", c.AppName).
		Str("version", c.AppVersion).
		Logger()

	if c.Pretty {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "15:04:05"})
	}

	log.Info().Msg("logger: initialized")
}
