package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Kbnh/tasks/config"
	"github.com/Kbnh/tasks/internal/adapter/postgres"
	"github.com/Kbnh/tasks/internal/controller/http"
	"github.com/Kbnh/tasks/internal/usecase"
	_ "github.com/Kbnh/tasks/pkg/logger" // Инициализируем логгер, импортируя пакет, который настраивает глобальный логгер при загрузке
	"github.com/rs/zerolog/log"

	"github.com/Kbnh/tasks/pkg/httpserver"
	pgpool "github.com/Kbnh/tasks/pkg/postgres"
	"github.com/Kbnh/tasks/pkg/router"
	"github.com/Kbnh/tasks/pkg/transaction"
)

func Run(ctx context.Context, c config.Config) error {
	pgPool, err := pgpool.New(ctx, c.Postgres) // Создаем новый пул соединений с базой данных PostgreSQL, используя настройки из конфигурации
	if err != nil {
		return fmt.Errorf("pgpool.New: %w", err)
	}

	transaction.Init(pgPool) // Инициализируем пакет transaction, передавая ему пул соединений, чтобы он мог использовать его для управления транзакциями

	uc := usecase.New(postgres.New()) // Создаем новый экземпляр usecase, передавая ему реализацию репозитория на основе PostgreSQL

	r := router.New()  // Создаем новый роутер для обработки HTTP-запросов
	http.Router(r, uc) // Регистрируем маршруты и обработчики HTTP, используя функцию Router из пакета http, передавая ей роутер и экземпляр usecase

	httpServer := httpserver.New(r, c.HTTP) // Создаем новый HTTP-сервер, передавая ему роутер и настройки из конфигурации

	log.Info().Msg(fmt.Sprintf("%s v%s: started", c.App.Name, c.App.Version)) // Логируем информацию о запуске приложения, включая его имя и версию из конфигурации

	sig := make(chan os.Signal, 1) // graceful shutdown
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	<-sig

	httpServer.Close() // закрываем HTTP-сервер, чтобы прекратить обработку новых запросов и начать процесс завершения работы приложения

	pgPool.Close() // закрываем пул соединений с базой данных, чтобы освободить ресурсы и завершить все активные соединения

	log.Info().Msg(fmt.Sprintf("%s v%s: stopped", c.App.Name, c.App.Version)) // Логируем информацию о завершении работы приложения, включая его имя и версию из конфигурации

	return nil
}
