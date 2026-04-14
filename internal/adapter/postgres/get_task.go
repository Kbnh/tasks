package postgres

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Kbnh/tasks/internal/domain"
	"github.com/Kbnh/tasks/pkg/transaction"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type GetTaskDTO struct { // Структура для хранения данных задачи, полученных из базы данных, с типами данных, соответствующими типам в PostgreSQL
	ID          pgtype.UUID
	Title       pgtype.Text
	Description pgtype.Text
	Completed   pgtype.Bool
	CreatedAt   pgtype.Timestamptz
	UpdatedAt   pgtype.Timestamptz
}

func (d *GetTaskDTO) ToDomain() domain.Task {
	var updatedAt *time.Time
	if d.UpdatedAt.Valid { // Проверяем, действительно ли поле UpdatedAt содержит валидное значение
		t := d.UpdatedAt.Time
		updatedAt = &t
	}

	return domain.Task{ // Преобразуем данные из GetTaskDTO в доменную модель Task, используя соответствующие поля и типы данных
		ID:          d.ID.Bytes,
		Title:       d.Title.String,
		Description: d.Description.String,
		Completed:   d.Completed.Bool,
		CreatedAt:   d.CreatedAt.Time,
		UpdatedAt:   updatedAt,
	}
}

func (d *GetTaskDTO) Dest() []any {
	return []any{ // Возвращаем слайс указателей на поля структуры GetTaskDTO для сканирования данных из базы данных
		&d.ID,
		&d.Title,
		&d.Description,
		&d.Completed,
		&d.CreatedAt,
		&d.UpdatedAt,
	}
}

func (p *Postgres) GetTask(ctx context.Context, id uuid.UUID) (domain.Task, error) {
	// SQL-запрос для получения задачи по ID, исключая удаленные задачи (где deleted_at IS NULL)
	const query = `SELECT id, title, description, completed, created_at, updated_at 
	FROM tasks 
	WHERE id = $1 AND deleted_at IS NULL`

	var task GetTaskDTO // Инициализируем структуру для хранения данных задачи, полученных из базы данных

	txOrPool := transaction.TryExtractTX(ctx) // Пытаемся извлечь текущую транзакцию из контекста, если она есть, иначе используем пул соединений

	err := txOrPool.QueryRow(ctx, query, id).Scan(task.Dest()...) // Выполняем SQL-запрос и сканируем результат в структуру GetTaskDTO
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) { // Если ошибка указывает на отсутствие строк, возвращаем ошибку domain.ErrNotFound
			return domain.Task{}, domain.ErrNotFound
		}
		return domain.Task{}, fmt.Errorf("txOrPool.QueryRow.Scan: %w", err)
	}

	return task.ToDomain(), nil // Преобразуем данные из GetTaskDTO в доменную модель Task и возвращаем ее
}
