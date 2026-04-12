package httpserver

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

type Config struct {
	Port string `default:"8080" envconfig:"HTTP_PORT"`
}

type Server struct {
	server *http.Server
}

func New(handler http.Handler, c Config) *Server {
	httpServer := &http.Server{
		Addr:         net.JoinHostPort("", c.Port),
		Handler:      handler,
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
	}

	s := &Server{
		server: httpServer,
	}

	go s.start()

	log.Info().Msg("http server: started on port: " + c.Port)

	return s
}

func (s *Server) start() {
	err := s.server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal().Err(err).Msg("http server: ListenAndServe")
	}
}

func (s *Server) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), 25*time.Second)
	defer cancel()

	err := s.server.Shutdown(ctx)
	if err != nil {
		log.Error().Err(err).Msg("http server: Shutdown")
	}

	log.Info().Msg("http server: shutdown gracefully")

}
