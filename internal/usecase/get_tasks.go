package usecase

import (
	"context"
	"fmt"

	"github.com/Kbnh/tasks/internal/dto"
	"github.com/rs/zerolog/log"
)

func (u *UseCase) GetTasks(ctx context.Context, req dto.GetTasksRequest) (dto.GetTasksResponse, error) {

	var res dto.GetTasksResponse // Инициализируем пустой ответ

	tasks, err := u.repo.GetTasks(ctx, req) // Вызываем метод репозитория для получения списка задач с учетом фильтров и пагинации
	if err != nil {
		return res, fmt.Errorf("repo.GetTasks: %w", err)
	}
	res.Tasks = tasks // Заполняем поле Tasks в ответе полученными задачами

	log.Info().Int("count", len(tasks)).Msg("Tasks retrieved successfully") // Логируем успешное получение задач с количеством полученных задач

	return res, nil // Возвращаем ответ с задачами и nil в качестве ошибки
}
