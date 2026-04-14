package dto

import "github.com/Kbnh/tasks/internal/domain"

type UpdateTaskRequest struct { // Структура запроса для обновления задачи, содержащая ID задачи и поля, которые могут быть обновлены (Title, Description, Completed)
	ID          string  `json:"id"`
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Completed   *bool   `json:"completed"`
	// IdempotencyKey string  `json:"idempotency_key"` // Поле для обеспечения идемпотентности запросов, позволяя клиенту повторять запрос без риска создания дубликатов или нежелательных изменений
}

func (r UpdateTaskRequest) Validate() error {
	if r.Title == nil && r.Description == nil && r.Completed == nil {
		return domain.ErrNoFieldsToUpdate
	}

	// if r.IdempotencyKey == "" { // Пока нет реализации редиса
	// 	return domain.ErrIndependencyKeyRequired
	// }

	return nil
}
