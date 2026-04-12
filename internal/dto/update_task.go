package dto

import "github.com/Kbnh/tasks/internal/domain"

type UpdateTaskRequest struct {
	ID          string  `json:"id"`
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Completed   *bool   `json:"completed"`
	// IdempotencyKey string  `json:"idempotency_key"`
}

func (r UpdateTaskRequest) Validate() error {
	if r.Title == nil && r.Description == nil && r.Completed == nil {
		return domain.ErrNoFieldsToUpdate
	}

	// if r.IdempotencyKey == "" {
	// 	return domain.ErrIndependencyKeyRequired
	// }

	return nil
}
