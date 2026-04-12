package dto

import "github.com/google/uuid"

type CreateTaskRequest struct {
	Title       string `json:"title" validate:"required,min=1,max=255"`
	Description string `json:"description"`
}

type CreateTaskResponse struct {
	ID uuid.UUID `json:"id"`
}
