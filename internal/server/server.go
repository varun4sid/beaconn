package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"fmt"
	"net/http"
	"time"

	"github.com/varun4sid/beaconn-backend/internal/api/auth"
	"github.com/varun4sid/beaconn-backend/internal/api/health"
	"github.com/varun4sid/beaconn-backend/internal/api/leave"
	"github.com/varun4sid/beaconn-backend/internal/config"
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
	s.Router.Mount("/leave", leave.Routes(cfg))
}

func (s *Server) Start() {
	fmt.Printf("Server started on port %d\n", s.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", s.Port), s.Router)
}
