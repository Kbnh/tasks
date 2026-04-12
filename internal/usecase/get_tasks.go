package usecase

import (
	"context"
	"fmt"

	"github.com/Kbnh/tasks/internal/dto"
)

func (u *UseCase) GetTasks(ctx context.Context, req dto.GetTasksRequest) (dto.GetTasksResponse, error) {
	var res dto.GetTasksResponse

	tasks, err := u.repo.GetTasks(ctx, req)
	if err != nil {
		return res, fmt.Errorf("repo.GetTasks: %w", err)
	}
	res.Tasks = tasks

	return res, nil
}
