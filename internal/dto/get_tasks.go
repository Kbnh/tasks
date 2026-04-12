package dto

import (
	"fmt"

	"github.com/Kbnh/tasks/internal/domain"
	"github.com/go-playground/validator/v10"
)

type GetTasksResponse struct {
	Tasks []domain.Task `json:"tasks"`
}

type GetTasksRequest struct {
	Sort  string `validate:"omitempty,oneofci=id title created_at"`
	Order string `validate:"omitempty,oneofci=asc desc"`
}

var validate = validator.New(validator.WithRequiredStructEnabled())

func (r GetTasksRequest) Validate() error {
	err := validate.Struct(r)
	if err != nil {
		return fmt.Errorf("validate.Struct: %w", err)
	}

	return nil
}
