package v1

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Kbnh/tasks/internal/domain"
	"github.com/Kbnh/tasks/internal/dto"
	"github.com/Kbnh/tasks/internal/usecase"
	"github.com/go-chi/chi/v5"
)

func updateTask(uc *usecase.UseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		var req dto.UpdateTaskRequest // создаем request dto

		idStr := chi.URLParam(r, "id") // получаем id задачи из query параметров
		if idStr == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		req.ID = idStr // устанавливаем id в request dto

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil { // декодируем json в структуру
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err := uc.UpdateTask(r.Context(), req) // вызываем usecase для обновления задачи
		if err != nil {
			if errors.Is(err, domain.ErrNotFound) { // если задача не найдена, возвращаем 404 Not Found
				w.WriteHeader(http.StatusNotFound)
				return
			}
			if errors.Is(err, domain.ErrNoFieldsToUpdate) { // если нет полей для обновления, возвращаем 400 Bad Request
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK) // устанавливаем статус код 200 OK
	}
}
