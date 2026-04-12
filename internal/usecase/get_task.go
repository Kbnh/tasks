package usecase

import (
	"context"
	"fmt"

	"github.com/Kbnh/tasks/internal/domain"
	"github.com/Kbnh/tasks/internal/dto"
	"github.com/google/uuid"
)

func (u *UseCase) GetTask(ctx context.Context, req dto.GetTaskRequest) (dto.GetTaskResponse, error) {
	var res dto.GetTaskResponse

	id, err := uuid.Parse(req.ID)
	if err != nil {
		return res, domain.ErrInvalidInput
	}

	task, err := u.repo.GetTask(ctx, id)
	if err != nil {
		return res, fmt.Errorf("repo.GetTask: %w", err)
	}

	if task.IsDeleted() {
		return res, domain.ErrNotFound
	}

	return dto.GetTaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Completed:   task.Completed,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}, nil
}
