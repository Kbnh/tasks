package config

import (
	"fmt"

	"github.com/Kbnh/tasks/pkg/httpserver"
	"github.com/Kbnh/tasks/pkg/logger"
	"github.com/Kbnh/tasks/pkg/postgres"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type App struct {
	Name    string `envconfig:"APP_NAME"    required:"true"`
	Version string `envconfig:"APP_VERSION" required:"true"`
}

type Config struct {
	App      App
	Postgres postgres.Config
	HTTP     httpserver.Config
	Logger   logger.Config
}

func New() (Config, error) {
	var c Config

	err := godotenv.Load(".env")
	if err != nil {
		return Config{}, fmt.Errorf("godotenv.Load: %w", err)
	}

	err = envconfig.Process("", &c)
	if err != nil {
		return Config{}, fmt.Errorf("envconfig.Process: %w", err)
	}

	return c, nil
}
