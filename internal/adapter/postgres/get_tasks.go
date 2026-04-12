package postgres

import (
	"context"
	"fmt"

	"github.com/Kbnh/tasks/internal/domain"
	"github.com/Kbnh/tasks/internal/dto"
	"github.com/Kbnh/tasks/pkg/transaction"
)

func (p *Postgres) GetTasks(ctx context.Context, req dto.GetTasksRequest) ([]domain.Task, error) {
	query := `SELECT id, title, description, completed, created_at, updated_at 
	FROM tasks 
	WHERE deleted_at IS NULL
	ORDER BY %s %s`

	sort := "created_at"
	if req.Sort != "" {
		sort = req.Sort
	}

	order := "ASC"
	if req.Order != "" {
		order = req.Order
	}

	query = fmt.Sprintf(query, sort, order)

	txOrPool := transaction.TryExtractTX(ctx)

	tasks := make([]domain.Task, 0)

	rows, err := txOrPool.Query(ctx, query)
	if err != nil {
		return tasks, fmt.Errorf("txOrPool.Query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var task GetTaskDTO
		err := rows.Scan(task.Dest()...)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan: %w", err)
		}
		tasks = append(tasks, task.ToDomain())
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows.Err: %w", err)
	}

	return tasks, nil
}
