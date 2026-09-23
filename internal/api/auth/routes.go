package auth

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	cfg "github.com/varun4sid/beaconn-backend/internal/config"
)

func Routes(c *cfg.Config) *chi.Mux {
	r := chi.NewRouter()

	r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		loginHandler(w, r, c)
	})

	return r
}
