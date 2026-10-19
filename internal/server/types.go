package server

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Server struct {
	Router    *chi.Mux
	Port      int
	Transport *http.Transport
}
