package httpserver

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

type Config struct { // Конфигурация для HTTP сервера, содержащая порт, на котором будет работать сервер
	Port string `default:"8080" envconfig:"HTTP_PORT"`
}

type Server struct { // Структура сервера, содержащая поле для хранения экземпляра http.Server
	server *http.Server
}

func New(handler http.Handler, c Config) *Server { // Конструктор для создания нового HTTP сервера

	httpServer := &http.Server{ // Создаем новый экземпляр http.Server с указанными параметрами, включая адрес, обработчик, таймауты для чтения и записи
		Addr:         net.JoinHostPort("", c.Port),
		Handler:      handler,
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
	}

	s := &Server{ // Создаем новый экземпляр Server, передавая ему созданный http.Server
		server: httpServer,
	}

	go s.start() // Запускаем сервер в отдельной горутине, чтобы он не блокировал выполнение основного потока приложения

	log.Info().Msg("http server: started on port: " + c.Port) // Логируем информацию о том, что сервер успешно запущен и на каком порту он работает

	return s
}

func (s *Server) start() { // Метод для запуска HTTP сервера, который вызывает метод ListenAndServe и обрабатывает возможные ошибки при запуске сервера
	err := s.server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed { // Проверяем, что ошибка не связана с нормальным завершением работы сервера (http.ErrServerClosed), и если это так, то логируем фатальную ошибку и завершаем приложение
		log.Fatal().Err(err).Msg("http server: ListenAndServe")
	}
}

func (s *Server) Close() { // Метод для graceful shutdown сервера, который создает контекст с таймаутом и вызывает метод Shutdown на http.Server, позволяя серверу завершить обработку текущих запросов перед остановкой
	ctx, cancel := context.WithTimeout(context.Background(), 25*time.Second) // Контекст
	defer cancel()                                                           // Отложенный вызов функции cancel для освобождения ресурсов контекста после завершения метода

	err := s.server.Shutdown(ctx) // Вызываем метод Shutdown на http.Server, передавая ему контекст с таймаутом, чтобы сервер мог завершить обработку текущих запросов и освободить ресурсы
	if err != nil {
		log.Error().Err(err).Msg("http server: Shutdown")
	}

	log.Info().Msg("http server: shutdown gracefully") // Логируем информацию о том, что сервер был успешно остановлен в режиме graceful shutdown

}
