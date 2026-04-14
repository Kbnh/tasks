package dto

import (
	"time"

	"github.com/google/uuid"
)

type GetTaskResponse struct { // Структура ответа для получения задачи, содержащая все поля задачи
	ID          uuid.UUID  `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	Title       string     `json:"title" validate:"required,min=1,max=255"`
	Description string     `json:"description"`
	Completed   bool       `json:"completed"`
	// DeletedAt не включаем в ответ, так как он используется только для внутренней логики и не должен быть видимым для клиентов API
}

type GetTaskRequest struct { // Структура запроса для получения задачи, содержащая ID задачи в виде строки
	ID string
}
