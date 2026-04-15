package usecase

import (
	"context"
	"fmt"

	"github.com/Kbnh/tasks/internal/domain"
	"github.com/Kbnh/tasks/internal/dto"
	"github.com/google/uuid"
)

func (u *UseCase) GetTask(ctx context.Context, req dto.GetTaskRequest) (dto.GetTaskResponse, error) {

	var res dto.GetTaskResponse // Инициализируем пустой ответ

	id, err := uuid.Parse(req.ID) // Парсим строковый ID в тип uuid.UUID
	if err != nil {
		return res, domain.ErrInvalidInput
	}

	task, err := u.repo.GetTask(ctx, id) // Вызываем метод репозитория для получения задачи по ID
	if err != nil {
		return res, fmt.Errorf("repo.GetTask: %w", err)
	}

	if task.IsDeleted() { // Проверяем, помечена ли задача как удаленная
		return res, domain.ErrNotFound
	}

	return dto.GetTaskResponse{ // Возвращаем данные задачи в ответе
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Completed:   task.Completed,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}, nil
}
