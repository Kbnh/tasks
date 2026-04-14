package transaction

import (
	"context"

	"github.com/Kbnh/tasks/pkg/postgres"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ( // Глобальная переменная для хранения пула соединений с базой данных PostgreSQL, которая будет использоваться для выполнения запросов к базе данных
	pool *pgxpool.Pool
)

type ctxKey struct{} // Пустая структура, которая будет использоваться в качестве ключа для хранения транзакции в контексте, чтобы избежать конфликтов с другими значениями в контексте

func Init(p *postgres.Pool) { // Функция для инициализации глобальной переменной pool
	pool = p.Pool
}

type Transaction struct { // Структура для хранения текущей транзакции
	pgx.Tx
}

type Executor interface { // Интерфейс для выполнения SQL-запросов, который может быть реализован как пулом соединений, так и транзакцией
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

func TryExtractTX(ctx context.Context) Executor { // Функция для извлечения текущей транзакции из контекста, если она есть, иначе возвращает пул соединений
	tx, ok := ctx.Value(ctxKey{}).(*Transaction) // Пытаемся извлечь значение из контекста по ключу ctxKey{} и привести его к типу *Transaction. Если это удается, возвращаем транзакцию, иначе возвращаем пул соединений
	if !ok {
		return pool
	}

	return tx
}
