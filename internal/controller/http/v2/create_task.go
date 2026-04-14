package v2

import (
	"encoding/json"
	"net/http"

	"github.com/Kbnh/tasks/internal/dto"
	"github.com/Kbnh/tasks/internal/usecase"
)

func createTask(uc *usecase.UseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		var req dto.CreateTaskRequest // создаем response dto

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil { // декодируем json в структуру
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		res, err := uc.CreateTaskV2(r.Context(), req) // вызываем usecase для создания задачи
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json") // устанавливаем заголовок Content-Type для ответа
		w.WriteHeader(http.StatusCreated)                  // устанавливаем статус код 201 Created
		json.NewEncoder(w).Encode(res)                     // кодируем результат в json и отправляем в ответ
	}
}
