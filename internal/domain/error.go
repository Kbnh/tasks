package domain

import "errors"

var ( // Кастомные ошибки
	ErrNotFound                = errors.New("not found")
	ErrAlreadyExists           = errors.New("already exists")
	ErrInvalidInput            = errors.New("invalid input")
	ErrNoFieldsToUpdate        = errors.New("no fields to update")
	ErrIndependencyKeyRequired = errors.New("idempotency key is required")
)
