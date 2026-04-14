package usecase

import (
	"context"
	"fmt"

	"github.com/Kbnh/tasks/internal/domain"
	"github.com/Kbnh/tasks/internal/dto"
	"github.com/Kbnh/tasks/pkg/transaction"
	"github.com/rs/zerolog/log"
)

func (u *UseCase) CreateTask(ctx context.Context, req dto.CreateTaskRequest) (dto.CreateTaskResponse, error) {

	var res dto.CreateTaskResponse // Инициализируем пустой ответ

	task, err := domain.NewTask(req.Title, req.Description) // Создаем новую задачу, используя данные из запроса
	if err != nil {
		return res, fmt.Errorf("domain.NewTask: %w", err)
	}

	err = transaction.Wrap(ctx, func(ctx context.Context) error { // Оборачиваем выполнение в транзакцию
		err = u.repo.CreateTask(ctx, task) // Сохраняем задачу в репозитории
		if err != nil {
			return fmt.Errorf("repo.CreateTask: %w", err)
		}

		return nil
	})
	if err != nil {
		return res, fmt.Errorf("transaction.Wrap: %w", err)
	}

	log.Info().Str("id", task.ID.String()).Msg("Task created successfully") // Логируем успешное создание задачи с ее ID

	return dto.CreateTaskResponse{ // Возвращаем ID созданной задачи в ответе
		ID: task.ID,
	}, nil
}
