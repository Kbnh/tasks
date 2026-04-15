package usecase

import (
	"context"
	"fmt"

	"github.com/Kbnh/tasks/internal/domain"
	"github.com/Kbnh/tasks/internal/dto"
	"github.com/Kbnh/tasks/pkg/transaction"
	"github.com/google/uuid"
)

func (u *UseCase) UpdateTask(ctx context.Context, req dto.UpdateTaskRequest) error {

	if err := req.Validate(); err != nil { // Валидируем входные данные запроса
		return fmt.Errorf("req.Validate: %w", err)
	}

	id, err := uuid.Parse(req.ID) // Парсим строковый ID в тип uuid.UUID
	if err != nil {
		return fmt.Errorf("uuid.Parse: %w", err)
	}

	err = transaction.Wrap(ctx, func(ctx context.Context) error { // Оборачиваем выполнение в транзакцию
		task, err := u.repo.GetTask(ctx, id) // Получаем текущую задачу из репозитория по ID
		if err != nil {
			return fmt.Errorf("repo.GetTask: %w", err)
		}

		if task.IsDeleted() { // Проверяем, помечена ли задача как удаленная
			return domain.ErrNotFound
		}

		newTask, ok := update(task, req) // Обновляем поля задачи на основе данных из запроса, возвращая флаг, указывающий, были ли изменения

		if !ok { // Если не было изменений, возвращаем ошибку, указывающую на отсутствие полей для обновления
			return domain.ErrNoFieldsToUpdate
		}

		err = u.repo.UpdateTask(ctx, &newTask) // Сохраняем обновленную задачу в репозитории
		if err != nil {
			return fmt.Errorf("repo.UpdateTask: %w", err)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("transaction.Wrap: %w", err)
	}

	return nil
}

func update(task domain.Task, req dto.UpdateTaskRequest) (domain.Task, bool) {
	ok := false // Инициализируем флаг, указывающий, были ли изменения
	if req.Title != nil {
		task.Title = *req.Title
		ok = true
	}
	if req.Description != nil {
		task.Description = *req.Description
		ok = true
	}
	if req.Completed != nil {
		task.Completed = *req.Completed
		ok = true
	}

	return task, ok
}
