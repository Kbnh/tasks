package domain

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Task struct { // Структура Task, представляющая задачу в системе
	ID          uuid.UUID  `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"-"`
	Title       string     `json:"title" validate:"required,min=1,max=255"`
	Description string     `json:"description"`
	Completed   bool       `json:"completed"`
}

var validate = validator.New(validator.WithRequiredStructEnabled())

func NewTask(title, description string) (*Task, error) { // Конструктор

	task := &Task{ // Создаем новый экземпляр Task
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: nil,
		DeletedAt: nil,

		Title:       title,
		Description: description,
		Completed:   false,
	}

	if err := task.Validate(); err != nil { // Валидируем созданную задачу, проверяя, что все обязательные поля заполнены и соответствуют требованиям
		return nil, fmt.Errorf("task.Validate: %w", err)
	}

	return task, nil
}

func (t *Task) Validate() error { // Метод для валидации полей задачи, используя пакет validator
	err := validate.Struct(t)
	if err != nil {
		return fmt.Errorf("validate.Struct: %w", err)
	}

	return nil
}

func (t *Task) IsDeleted() bool { // Метод для проверки, помечена ли задача как удаленная
	return t.DeletedAt != nil
}
