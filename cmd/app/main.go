package main

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/Kbnh/tasks/config"
	"github.com/Kbnh/tasks/internal/app"
	"github.com/Kbnh/tasks/pkg/logger" // Инициализируем логгер, импортируя пакет, который настраивает глобальный логгер при загрузке
)

func main() { // Точка входа в приложение, которая выполняет следующие шаги:
	c, err := config.New() // Загружает конфигурацию приложения
	if err != nil {
		log.Fatal().Err(err).Msg("config.New")
	}

	closer, err := logger.Init(c.Logger)
	if err != nil {
		log.Fatal().Err(err).Msg("logger.Init")
	}
	defer closer.Close()

	ctx := context.Background() // Создает новый контекст

	err = app.Run(ctx, c) // Запускает приложение, передавая ему контекст и конфигурацию
	if err != nil {
		log.Fatal().Err(err).Msg("app.Run")
	}
}
