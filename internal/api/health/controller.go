package health

import (
	"net/http"

	"github.com/varun4sid/beaconn-backend/internal/api"
)

type HealthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{
		Status:  "success",
		Message: "API is healthy",
	}

	api.WriteResponseBody(w, response)
}
