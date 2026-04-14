package v1

import (
	"encoding/json"
	"net/http"

	"github.com/Kbnh/tasks/internal/dto"
	"github.com/Kbnh/tasks/internal/usecase"
	"github.com/go-chi/chi/v5"
)

func getTask(uc *usecase.UseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req dto.GetTaskRequest // создаем request dto

		idStr := chi.URLParam(r, "id") // получаем id задачи из query параметров
		if idStr == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		req.ID = idStr // устанавливаем id в request dto

		task, err := uc.GetTask(r.Context(), req) // вызываем usecase для получения задачи
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)    // устанавливаем статус код 200 OK
		json.NewEncoder(w).Encode(task) // кодируем задачу в json и отправляем в ответе
	}
}
