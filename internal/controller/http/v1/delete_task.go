package v1

import (
	"net/http"

	"github.com/Kbnh/tasks/internal/dto"
	"github.com/Kbnh/tasks/internal/usecase"
	"github.com/go-chi/chi/v5"
)

func deleteTask(uc *usecase.UseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req dto.DeleteTaskRequest // создаем request dto

		idStr := chi.URLParam(r, "id") // получаем id задачи из query параметров
		if idStr == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		req.ID = idStr

		err := uc.DeleteTask(r.Context(), req) // вызываем usecase для удаления задачи
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent) // устанавливаем статус код 204 No Content
	}
}
