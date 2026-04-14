package postgres

import (
	"context"
	"fmt"

	"github.com/Kbnh/tasks/internal/domain"
	"github.com/Kbnh/tasks/pkg/transaction"
)

func (p *Postgres) CreateTask(ctx context.Context, task *domain.Task) error {
	// SQL-запрос для вставки новой задачи в таблицу tasks
	const query = `
		INSERT INTO tasks (id, created_at, updated_at, deleted_at, title, description, completed)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	txOrPool := transaction.TryExtractTX(ctx) // Пытаемся извлечь текущую транзакцию из контекста, если она есть, иначе используем пул соединений

	// Выполняем SQL-запрос с параметрами, соответствующими полям задачи
	_, err := txOrPool.Exec(ctx, query, task.ID, task.CreatedAt, task.UpdatedAt, task.DeletedAt, task.Title, task.Description, task.Completed)
	if err != nil {
		return fmt.Errorf("txOrPool.Exec: %w", err)
	}

	return nil
}
