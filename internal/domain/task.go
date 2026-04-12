package domain

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID  `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"-"`
	Title       string     `json:"title" validate:"required,min=1,max=255"`
	Description string     `json:"description"`
	Completed   bool       `json:"completed"`
}

var validate = validator.New(validator.WithRequiredStructEnabled())

func NewTask(title, description string) (*Task, error) {
	task := &Task{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: nil,
		DeletedAt: nil,

		Title:       title,
		Description: description,
		Completed:   false,
	}

	if err := task.Validate(); err != nil {
		return nil, fmt.Errorf("task.Validate: %w", err)
	}

	return task, nil
}

func (t *Task) Validate() error {
	err := validate.Struct(t)
	if err != nil {
		return fmt.Errorf("validate.Struct: %w", err)
	}

	return nil
}

func (t *Task) IsDeleted() bool {
	return t.DeletedAt != nil
}
