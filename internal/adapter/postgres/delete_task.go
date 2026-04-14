package postgres

import (
	"context"
	"fmt"

	"github.com/Kbnh/tasks/pkg/transaction"
	"github.com/google/uuid"
)

func (p *Postgres) DeleteTask(ctx context.Context, id uuid.UUID) error {
	// SQL-запрос для логического удаления задачи, устанавливая поле deleted_at в текущее время, если задача не была удалена ранее
	const query = `UPDATE tasks SET deleted_at = NOW()
	WHERE id = $1 AND deleted_at IS NULL`

	txOrPool := transaction.TryExtractTX(ctx) // Пытаемся извлечь текущую транзакцию из контекста, если она есть, иначе используем пул соединений

	_, err := txOrPool.Exec(ctx, query, id) // Выполняем SQL-запрос с параметром ID задачи для логического удаления
	if err != nil {
		return fmt.Errorf("txOrPool.Exec: %w", err)
	}

	return nil
}
