package server

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/varun4sid/beaconn-backend/internal/api/auth"
	"github.com/varun4sid/beaconn-backend/internal/api/health"
	"github.com/varun4sid/beaconn-backend/internal/api/leave"
	"github.com/varun4sid/beaconn-backend/internal/api/token"
	"github.com/varun4sid/beaconn-backend/internal/config"
)

type Server struct {
	Router    *chi.Mux
	Port      int
	Transport *http.Transport
}

func NewServer(cfg *config.Config) *Server {
	s := &Server{
		Router:    chi.NewRouter(),
		Port:      cfg.Port,
		Transport: cfg.Transport,
	}

	s.middleware(cfg)
	s.routes(cfg)

	return s
}

func (s *Server) middleware(cfg *config.Config) {
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Recoverer)
	s.Router.Use(middleware.Timeout(60 * time.Second))
	s.Router.Use(corsMiddleware(cfg.AppEnv, cfg.AllowedDomain))
}

func corsMiddleware(appEnv, allowedDomain string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")
			if originAllowed(appEnv, allowedDomain, origin) {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
				w.Header().Add("Vary", "Origin")
			}
			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func originAllowed(appEnv, allowedDomain, origin string) bool {
	if origin == "" {
		return false
	}

	if appEnv == "development" &&
		(origin == "http://localhost:4200" || origin == "http://127.0.0.1:4200") {
		return true
	}

	if allowedDomain == origin {
		return true
	}

	if !strings.Contains(allowedDomain, "://") {
		return origin == fmt.Sprintf("http://%s", allowedDomain) ||
			origin == fmt.Sprintf("https://%s", allowedDomain)
	}

	return false
}

func (s *Server) routes(cfg *config.Config) {
	s.Router.Mount("/health", health.Routes())
	s.Router.Mount("/auth", auth.Routes(cfg))
	s.Router.Mount("/leave", leave.Routes(cfg))
	s.Router.Mount("/token", token.Routes(cfg))
}

func (s *Server) Start() {
	fmt.Printf("Server started on port %d\n", s.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", s.Port), s.Router)
}
