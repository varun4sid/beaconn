package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"net/http"
	"time"
	"fmt"

	"github.com/varun4sid/beaconn-backend/internal/api/health"
	"github.com/varun4sid/beaconn-backend/internal/config"
	"github.com/varun4sid/beaconn-backend/internal/api/auth"
)

type Server struct {
	Router *chi.Mux
	Port   int
	Transport *http.Transport
}

func NewServer(cfg *config.Config) *Server {
	s := &Server{
		Router: chi.NewRouter(),
		Port:   cfg.Port,
		Transport: cfg.Transport,
	}

	s.middleware()
	s.routes(cfg)

	return s
}

func (s *Server) middleware() {
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Recoverer)
	s.Router.Use(middleware.Timeout(60 * time.Second))
}

func (s *Server) routes(cfg *config.Config) {
	s.Router.Mount("/health", health.Routes())
	s.Router.Mount("/auth", auth.Routes(cfg))
}

func (s *Server) Start() {
	fmt.Printf("Server started on port %d\n", s.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", s.Port), s.Router)
}
