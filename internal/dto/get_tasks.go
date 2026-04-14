package dto

import (
	"fmt"

	"github.com/Kbnh/tasks/internal/domain"
	"github.com/go-playground/validator/v10"
)

type GetTasksResponse struct { // Структура ответа для получения списка задач, содержащая слайс задач
	Tasks []domain.Task `json:"tasks"`
}

type GetTasksRequest struct { // Структура запроса для получения списка задач, содержащая параметры для фильтрации, сортировки и пагинации
	Sort  string `validate:"omitempty,oneofci=id title created_at"`
	Order string `validate:"omitempty,oneofci=asc desc"`
}

var validate = validator.New(validator.WithRequiredStructEnabled())

func (r GetTasksRequest) Validate() error { // Метод для валидации полей запроса, проверяя, что указанные поля для сортировки и направления сортировки соответствуют допустимым значениям
	err := validate.Struct(r)
	if err != nil {
		return fmt.Errorf("validate.Struct: %w", err)
	}

	return nil
}
