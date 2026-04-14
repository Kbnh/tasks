package transaction

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
)

func Wrap(ctx context.Context, fn func(ctx context.Context) error) error { // Функция для оборачивания выполнения в транзакцию
	tx, err := pool.Begin(ctx) // Начинаем новую транзакцию, используя пул соединений
	if err != nil {
		return fmt.Errorf("pool.Begin: %w", err)
	}

	defer func() { // Отложенная функция для отката транзакции в случае ошибки
		err := tx.Rollback(ctx)                             // Откатываем транзакцию, если она еще не была коммитирована
		if err != nil && !errors.Is(err, pgx.ErrTxClosed) { // Если произошла ошибка при откате, и эта ошибка не связана с тем, что транзакция уже была закрыта (коммитирована), то логируем эту ошибку
			log.Error().Err(err).Msg("transaction.Rollback")
		}
	}()

	ctx = context.WithValue(ctx, ctxKey{}, &Transaction{tx}) // Создаем новый контекст, в который помещаем текущую транзакцию

	err = fn(ctx) // Вызываем переданную функцию с новым контекстом, который содержит текущую транзакцию
	if err != nil {
		return fmt.Errorf("fn: %w", err)
	}

	err = tx.Commit(ctx) // Коммитим транзакцию, если функция выполнилась успешно
	if err != nil {
		return fmt.Errorf("tx.Commit: %w", err)
	}

	return nil
}
