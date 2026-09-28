package token

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	cfg "github.com/varun4sid/beaconn-backend/internal/config"
)

func Routes(c *cfg.Config) *chi.Mux {
	r := chi.NewRouter()

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		tokenHandler(w, r, c)
	})

	return r
}
