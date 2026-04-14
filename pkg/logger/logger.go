package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Config struct { // Конфигурация для логгера
	AppName    string `envconfig:"APP_NAME" required:"true"`
	AppVersion string `envconfig:"APP_VERSION" required:"true"`
	Level      string `default:"error" envconfig:"LOGGER_LEVEL"`
	Pretty     bool   `default:"false" envconfig:"LOGGER_PRETTY"`
}

func Init(c Config) { // Функция для инициализации логгера, которая настраивает формат времени, уровень логирования и добавляет общие поля для всех логов, такие как имя приложения и версия
	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	if level, err := zerolog.ParseLevel(c.Level); err == nil && level != zerolog.NoLevel { // Парсим уровень логирования из конфигурации и устанавливаем его, если он валиден и не равен zerolog.NoLevel
		zerolog.SetGlobalLevel(level)
	}

	log.Logger = log.With(). // Добавляем общие поля для всех логов, такие как имя приложения и версия, чтобы упростить фильтрацию и анализ логов в будущем
					Str("app", c.AppName).
					Str("version", c.AppVersion).
					Logger()

	if c.Pretty { // Если в конфигурации указано, что логирование должно быть в "красивом" формате, то настраиваем логгер на вывод в консоль с форматированием, что может быть полезно для разработки и отладки
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "15:04:05"})
	}

	log.Info().Msg("logger: initialized") // Логируем информацию о том, что логгер был успешно инициализирован
}
