package v1

import (
	"net/http"

	"github.com/Kbnh/tasks/internal/usecase"
)

func CreateTask(uc *usecase.UseCase) http.HandlerFunc {
	return createTask(uc)
}

func GetTasks(uc *usecase.UseCase) http.HandlerFunc {
	return getTasks(uc)
}

func GetTask(uc *usecase.UseCase) http.HandlerFunc {
	return getTask(uc)
}

func UpdateTask(uc *usecase.UseCase) http.HandlerFunc {
	return updateTask(uc)
}

func DeleteTask(uc *usecase.UseCase) http.HandlerFunc {
	return deleteTask(uc)
}
