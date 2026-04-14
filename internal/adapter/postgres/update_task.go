package postgres

import (
	"context"
	"fmt"

	"github.com/Kbnh/tasks/internal/domain"
	"github.com/Kbnh/tasks/pkg/transaction"
)

func (p *Postgres) UpdateTask(ctx context.Context, task *domain.Task) error {
	// SQL-запрос для обновления существующей задачи, устанавливая новые значения для полей title, description, completed и обновляя поле updated_at, если задача не была удалена ранее
	const query = `
		UPDATE tasks
		SET title = $1, description = $2, completed = $3, updated_at = NOW()
		WHERE id = $4 AND deleted_at IS NULL
	`

	txOrPool := transaction.TryExtractTX(ctx) // Пытаемся извлечь текущую транзакцию из контекста, если она есть, иначе используем пул соединений

	result, err := txOrPool.Exec(ctx, query, task.Title, task.Description, task.Completed, task.ID) // Выполняем SQL-запрос с параметрами, соответствующими полям задачи для обновления
	if err != nil {
		return fmt.Errorf("txOrPool.Exec: %w", err)
	}

	if result.RowsAffected() == 0 { // Если количество затронутых строк равно нулю, это означает, что задача с указанным ID не была найдена или уже была удалена, поэтому возвращаем ошибку domain.ErrNotFound
		return domain.ErrNotFound
	}

	return nil
}
