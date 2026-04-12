package v1

import (
	"encoding/json"
	"net/http"

	"github.com/Kbnh/tasks/internal/dto"
	"github.com/Kbnh/tasks/internal/usecase"
)

func getTasks(uc *usecase.UseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		req := dto.GetTasksRequest{ // создаем request dto, заполняя его данными из query параметров
			Sort:  r.URL.Query().Get("sort"),
			Order: r.URL.Query().Get("order"),
		}

		if req.Validate() != nil { // валидируем request dto
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		tasks, err := uc.GetTasks(r.Context(), req) // вызываем usecase для получения всех задач
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)     // устанавливаем статус код 200 OK
		json.NewEncoder(w).Encode(tasks) // кодируем задачи в json и отправляем в ответе
	}
}
