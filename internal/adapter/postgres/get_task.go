package postgres

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Kbnh/tasks/internal/domain"
	"github.com/Kbnh/tasks/pkg/transaction"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type GetTaskDTO struct {
	ID          pgtype.UUID
	Title       pgtype.Text
	Description pgtype.Text
	Completed   pgtype.Bool
	CreatedAt   pgtype.Timestamptz
	UpdatedAt   pgtype.Timestamptz
}

func (d *GetTaskDTO) ToDomain() domain.Task {
	var updatedAt *time.Time
	if d.UpdatedAt.Valid {
		t := d.UpdatedAt.Time
		updatedAt = &t
	}

	return domain.Task{
		ID:          d.ID.Bytes,
		Title:       d.Title.String,
		Description: d.Description.String,
		Completed:   d.Completed.Bool,
		CreatedAt:   d.CreatedAt.Time,
		UpdatedAt:   updatedAt,
	}
}

func (d *GetTaskDTO) Dest() []any {
	return []any{
		&d.ID,
		&d.Title,
		&d.Description,
		&d.Completed,
		&d.CreatedAt,
		&d.UpdatedAt,
	}
}

func (p *Postgres) GetTask(ctx context.Context, id uuid.UUID) (domain.Task, error) {
	const query = `SELECT id, title, description, completed, created_at, updated_at 
	FROM tasks 
	WHERE id = $1 AND deleted_at IS NULL`

	var task GetTaskDTO

	txOrPool := transaction.TryExtractTX(ctx)

	err := txOrPool.QueryRow(ctx, query, id).Scan(task.Dest()...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Task{}, domain.ErrNotFound
		}
		return domain.Task{}, fmt.Errorf("txOrPool.QueryRow.Scan: %w", err)
	}

	return task.ToDomain(), nil
}
