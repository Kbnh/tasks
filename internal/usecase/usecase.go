package usecase

import (
	"context"

	"github.com/Kbnh/tasks/internal/domain"
	"github.com/Kbnh/tasks/internal/dto"
	"github.com/google/uuid"
)

type Repo interface { // Интерфейс репозитория, определяющий методы для работы с задачами
	CreateTask(ctx context.Context, task *domain.Task) error
	DeleteTask(ctx context.Context, id uuid.UUID) error
	GetTask(ctx context.Context, id uuid.UUID) (domain.Task, error)
	GetTasks(ctx context.Context, req dto.GetTasksRequest) ([]domain.Task, error)
	UpdateTask(ctx context.Context, task *domain.Task) error
}

type UseCase struct { // Структура UseCase, которая содержит ссылку на репозиторий для выполнения бизнес-логики
	repo Repo
}

func New(repo Repo) *UseCase { // Конструктор для создания нового экземпляра UseCase с переданным репозиторием
	return &UseCase{
		repo: repo,
	}
}
