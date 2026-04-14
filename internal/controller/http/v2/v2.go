package v2

import (
	"net/http"

	"github.com/Kbnh/tasks/internal/usecase"
)

func CreateTask(uc *usecase.UseCase) http.HandlerFunc {
	return createTask(uc)
}
