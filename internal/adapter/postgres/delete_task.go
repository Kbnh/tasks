package postgres

import (
	"context"
	"fmt"

	"github.com/Kbnh/tasks/pkg/transaction"
	"github.com/google/uuid"
)

func (p *Postgres) DeleteTask(ctx context.Context, id uuid.UUID) error {
	const query = `UPDATE tasks SET deleted_at = NOW() WHERE id = $1 AND deleted_at IS NULL`

	txOrPool := transaction.TryExtractTX(ctx)

	_, err := txOrPool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("txOrPool.Exec: %w", err)
	}

	return nil
}
