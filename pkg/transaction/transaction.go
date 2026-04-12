package transaction

import (
	"context"

	"github.com/Kbnh/tasks/pkg/postgres"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	pool *pgxpool.Pool
)

type ctxKey struct{}

func Init(p *postgres.Pool) {
	pool = p.Pool
}

type Transaction struct {
	pgx.Tx
}

type Executor interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

func TryExtractTX(ctx context.Context) Executor {
	tx, ok := ctx.Value(ctxKey{}).(*Transaction)
	if !ok {
		return pool
	}

	return tx
}
