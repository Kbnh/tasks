package dto

import "github.com/google/uuid"

type CreateTaskRequest struct { // Структура запроса для создания задачи, содержащая заголовок и описание задачи
	Title       string `json:"title" validate:"required,min=1,max=255"`
	Description string `json:"description"`
}

type CreateTaskResponse struct { // Структура ответа для создания задачи, содержащая ID созданной задачи
	ID uuid.UUID `json:"id"`
}
