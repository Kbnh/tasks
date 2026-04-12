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
	"github.com/rs/zerolog/log"

	"github.com/Kbnh/tasks/pkg/httpserver"
	pgpool "github.com/Kbnh/tasks/pkg/postgres"
	"github.com/Kbnh/tasks/pkg/router"
	"github.com/Kbnh/tasks/pkg/transaction"
)

func Run(ctx context.Context, c config.Config) error {
	pgPool, err := pgpool.New(ctx, c.Postgres)
	if err != nil {
		return fmt.Errorf("pgpool.New: %w", err)
	}

	transaction.Init(pgPool)

	uc := usecase.New(postgres.New())

	r := router.New()
	http.Router(r, uc)

	httpServer := httpserver.New(r, c.HTTP)

	log.Info().Msg(fmt.Sprintf("%s v%s: started", c.App.Name, c.App.Version))

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	<-sig

	httpServer.Close()

	pgPool.Close()

	log.Info().Msg(fmt.Sprintf("%s v%s: stopped", c.App.Name, c.App.Version))

	return nil
}
