package postgres

import (
	"context"
	"fmt"

	"github.com/Kbnh/tasks/internal/domain"
	"github.com/Kbnh/tasks/pkg/transaction"
)

func (p *Postgres) UpdateTask(ctx context.Context, task *domain.Task) error {
	const query = `
		UPDATE tasks
		SET title = $1, description = $2, completed = $3, updated_at = NOW()
		WHERE id = $4 AND deleted_at IS NULL
	`

	txOrPool := transaction.TryExtractTX(ctx)

	result, err := txOrPool.Exec(ctx, query, task.Title, task.Description, task.Completed, task.ID)
	if err != nil {
		return fmt.Errorf("txOrPool.Exec: %w", err)
	}

	if result.RowsAffected() == 0 {
		return domain.ErrNotFound
	}

	return nil
}
