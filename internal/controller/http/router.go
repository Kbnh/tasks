package http

import (
	v1 "github.com/Kbnh/tasks/internal/controller/http/v1"
	v2 "github.com/Kbnh/tasks/internal/controller/http/v2"
	"github.com/Kbnh/tasks/internal/usecase"
	"github.com/go-chi/chi/v5"
)

func Router(r chi.Router, uc *usecase.UseCase) { // Функция Router регистрирует маршруты и соответствующие обработчики для HTTP-запросов, используя роутер chi и экземпляр usecase для обработки бизнес-логики
	// Task v1 endpoints
	r.Post("/api/v1/tasks", v1.CreateTask(uc))
	r.Get("/api/v1/tasks", v1.GetTasks(uc))
	r.Get("/api/v1/tasks/{id}", v1.GetTask(uc))
	r.Patch("/api/v1/tasks/{id}", v1.UpdateTask(uc))
	r.Delete("/api/v1/tasks/{id}", v1.DeleteTask(uc))

	// Task v2 endpoints (добавлю позже POST, будет возвращать помимо id еще title и description)
	r.Post("/api/v2/tasks", v2.CreateTask(uc))
}
