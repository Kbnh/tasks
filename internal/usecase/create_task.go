package usecase

import (
	"context"
	"fmt"

	"github.com/Kbnh/tasks/internal/domain"
	"github.com/Kbnh/tasks/internal/dto"
	"github.com/Kbnh/tasks/pkg/transaction"
)

func (u *UseCase) CreateTask(ctx context.Context, req dto.CreateTaskRequest) (dto.CreateTaskResponse, error) {
	var res dto.CreateTaskResponse

	task, err := domain.NewTask(req.Title, req.Description)
	if err != nil {
		return res, fmt.Errorf("domain.NewTask: %w", err)
	}

	err = transaction.Wrap(ctx, func(ctx context.Context) error {
		err = u.repo.CreateTask(ctx, task)
		if err != nil {
			return fmt.Errorf("repo.CreateTask: %w", err)
		}

		return nil
	})
	if err != nil {
		return res, fmt.Errorf("transaction.Wrap: %w", err)
	}

	return dto.CreateTaskResponse{
		ID: task.ID,
	}, nil
}
