package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"net/http"
	"time"
	"fmt"

	"github.com/varun4sid/beaconn-backend/internal/api/health"
	"github.com/varun4sid/beaconn-backend/internal/config"
)

type Server struct {
	Router *chi.Mux
	Port   int
}

func NewServer(cfg *config.Config) *Server {
	s := &Server{
		Router: chi.NewRouter(),
		Port:   cfg.Port,
	}

	s.middleware()
	s.routes()

	return s
}

func (s *Server) middleware() {
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Recoverer)
	s.Router.Use(middleware.Timeout(60 * time.Second))
}

func (s *Server) routes() {
	s.Router.Mount("/health", health.Routes())
}

func (s *Server) Start() {	
	http.ListenAndServe(fmt.Sprintf(":%d", s.Port), s.Router)
}