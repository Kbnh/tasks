package usecase

import (
	"context"

	"github.com/Kbnh/tasks/internal/domain"
	"github.com/Kbnh/tasks/internal/dto"
	"github.com/google/uuid"
)

type Repo interface {
	CreateTask(ctx context.Context, task *domain.Task) error
	DeleteTask(ctx context.Context, id uuid.UUID) error
	GetTask(ctx context.Context, id uuid.UUID) (domain.Task, error)
	GetTasks(ctx context.Context, req dto.GetTasksRequest) ([]domain.Task, error)
	UpdateTask(ctx context.Context, task *domain.Task) error
}

type UseCase struct {
	repo Repo
}

func New(repo Repo) *UseCase {
	return &UseCase{
		repo: repo,
	}
}
