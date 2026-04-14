package config

import (
	"fmt"

	"github.com/Kbnh/tasks/pkg/httpserver"
	"github.com/Kbnh/tasks/pkg/logger"
	"github.com/Kbnh/tasks/pkg/postgres"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type App struct { // Конфигурация приложения
	Name    string `envconfig:"APP_NAME"    required:"true"`
	Version string `envconfig:"APP_VERSION" required:"true"`
}

type Config struct { // Общий конфиг
	App      App
	Postgres postgres.Config
	HTTP     httpserver.Config
	Logger   logger.Config
}

func New() (Config, error) { // Конструктор для создания новой конфигурации

	var c Config // Создаем переменную типа Config для хранения конфигурации приложения

	err := godotenv.Load(".env") // Загружаем переменные окружения из файла .env
	if err != nil {
		return Config{}, fmt.Errorf("godotenv.Load: %w", err)
	}

	err = envconfig.Process("", &c) // Заполняем структуру Config данными из переменных окружения
	if err != nil {
		return Config{}, fmt.Errorf("envconfig.Process: %w", err)
	}

	return c, nil // Возвращаем заполненную конфигурацию и nil в качестве ошибки
}
