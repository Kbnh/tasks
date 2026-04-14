package dto

import "github.com/google/uuid"

type CreateTaskResponseV2 struct { // Структура ответа для создания задачи, содержащая ID созданной задачи
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title" validate:"required,min=1,max=255"`
	Description string    `json:"description"`
}
