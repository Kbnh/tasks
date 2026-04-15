package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func New(middlewares ...func(http.Handler) http.Handler) chi.Router { // Функция для создания нового маршрутизатора, которая настраивает маршруты для проверки живости и готовности сервера
	r := chi.NewRouter() // Создаем новый маршрутизатор

	for _, mw := range middlewares {
		r.Use(mw)
	}

	r.Get("/live", probe)  // Регистрируем маршрут для проверки живости сервера, который будет обрабатывать GET-запросы на путь "/live" и вызывать функцию probe для обработки этих запросов
	r.Get("/ready", probe) // Регистрируем маршрут для проверки готовности сервера, который будет обрабатывать GET-запросы на путь "/ready" и вызывать функцию probe для обработки этих запросов

	return r
}

func probe(w http.ResponseWriter, r *http.Request) { // Функция-обработчик для маршрутов проверки живости и готовности сервера, которая просто возвращает статус 204 No Content
	w.WriteHeader(http.StatusNoContent)
}
