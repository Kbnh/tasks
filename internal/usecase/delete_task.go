package usecase

import (
	"context"
	"fmt"

	"github.com/Kbnh/tasks/internal/domain"
	"github.com/Kbnh/tasks/internal/dto"
	"github.com/google/uuid"
)

func (u *UseCase) DeleteTask(ctx context.Context, req dto.DeleteTaskRequest) error {

	id, err := uuid.Parse(req.ID) // Парсим строковый ID в тип uuid.UUID
	if err != nil {
		return domain.ErrInvalidInput
	}

	err = u.repo.DeleteTask(ctx, id) // Вызываем метод репозитория для удаления задачи по ID
	if err != nil {
		return fmt.Errorf("repo.DeleteTask: %w", err)
	}

	return nil

}
