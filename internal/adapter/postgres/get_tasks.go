package postgres

import (
	"context"
	"fmt"

	"github.com/Kbnh/tasks/internal/domain"
	"github.com/Kbnh/tasks/internal/dto"
	"github.com/Kbnh/tasks/pkg/transaction"
)

func (p *Postgres) GetTasks(ctx context.Context, req dto.GetTasksRequest) ([]domain.Task, error) {
	// SQL-запрос для получения списка задач, исключая удаленные задачи (где deleted_at IS NULL), с возможностью сортировки и указания порядка сортировки
	query := `SELECT id, title, description, completed, created_at, updated_at 
	FROM tasks 
	WHERE deleted_at IS NULL
	ORDER BY %s %s`

	sort := "created_at"
	if req.Sort != "" { // Если в запросе указано поле для сортировки, используем его вместо значения по умолчанию
		sort = req.Sort
	}

	order := "ASC"
	if req.Order != "" { // Если в запросе указано направление сортировки, используем его вместо значения по умолчанию
		order = req.Order
	}

	query = fmt.Sprintf(query, sort, order) // Форматируем SQL-запрос, вставляя выбранные поля для сортировки и направления сортировки

	txOrPool := transaction.TryExtractTX(ctx) // Пытаемся извлечь текущую транзакцию из контекста, если она есть, иначе используем пул соединений

	tasks := make([]domain.Task, 0) // Инициализируем слайс для хранения задач, которые будут получены из базы данных

	rows, err := txOrPool.Query(ctx, query) // Выполняем SQL-запрос для получения списка задач с учетом сортировки и направления сортировки
	if err != nil {
		return tasks, fmt.Errorf("txOrPool.Query: %w", err)
	}
	defer rows.Close()

	for rows.Next() { // Итерируемся по результатам запроса, сканируя каждую строку в структуру GetTaskDTO и преобразуя ее в доменную модель Task, которую добавляем в слайс tasks
		var task GetTaskDTO
		err := rows.Scan(task.Dest()...)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan: %w", err)
		}
		tasks = append(tasks, task.ToDomain())
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows.Err: %w", err)
	}

	return tasks, nil
}
