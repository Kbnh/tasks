package logger

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Config struct { // Конфигурация для логгера
	AppName    string `envconfig:"APP_NAME"    required:"true"`
	AppVersion string `envconfig:"APP_VERSION" required:"true"`
	Level      string `default:"error"         envconfig:"LOGGER_LEVEL"`
	Pretty     bool   `default:"false"         envconfig:"LOGGER_PRETTY"`
	Folder     string `default:"out/logs/"     envconfig:"LOGGER_FOLDER"`
}

func Init(c Config) (io.Closer, error) {
	zerolog.TimeFieldFormat = time.RFC3339
	if level, err := zerolog.ParseLevel(c.Level); err == nil && level != zerolog.NoLevel {
		zerolog.SetGlobalLevel(level)
	}

	if err := os.MkdirAll(c.Folder, 0755); err != nil {
		return nil, fmt.Errorf("os.MkdirAll: %w", err)
	}

	timestamp := time.Now().UTC().Format("2006-01-02T15-04-05.000000")
	logFilePath := filepath.Join(c.Folder, fmt.Sprintf("%s.log", timestamp))

	file, err := os.Create(logFilePath)
	if err != nil {
		return nil, fmt.Errorf("os.Create: %w", err)
	}

	var writers []io.Writer

	if c.Pretty {
		writers = append(writers, zerolog.ConsoleWriter{
			Out:        os.Stderr,
			TimeFormat: "15:04:05",
		})
	} else {
		writers = append(writers, os.Stderr)
	}

	writers = append(writers, file)

	multi := zerolog.MultiLevelWriter(writers...)

	log.Logger = log.With().
		// Caller().
		Str("app_name", c.AppName).
		Str("app_version", c.AppVersion).
		Logger().
		Output(multi)

	log.Info().Msg("logger initialized")

	return file, nil
}
