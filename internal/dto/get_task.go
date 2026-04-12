package dto

import (
	"time"

	"github.com/google/uuid"
)

type GetTaskResponse struct {
	ID          uuid.UUID  `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	Title       string     `json:"title" validate:"required,min=1,max=255"`
	Description string     `json:"description"`
	Completed   bool       `json:"completed"`
}

type GetTaskRequest struct {
	ID string
}
