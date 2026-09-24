package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/varun4sid/beaconn-backend/internal/api"
	cfg "github.com/varun4sid/beaconn-backend/internal/config"
	"github.com/varun4sid/beaconn-backend/internal/scraper"
)

func loginHandler(w http.ResponseWriter, r *http.Request, c *cfg.Config) {
	body, err := api.ParseRequestBody[AuthRequest](r)
	if err != nil {
		api.HttpError(w, fmt.Errorf("couldn't parse login request: %w", err), http.StatusBadRequest)
		return
	}

	if body.Rollno == "" || body.Password == "" {
		api.HttpError(w, fmt.Errorf("rollno and password are required"), http.StatusBadRequest)
		return
	}

	rollno := strings.ToUpper(body.Rollno)

	s, err := scraper.NewSession(c)
	if err != nil {
		api.HttpError(w, fmt.Errorf("failed to create scraper session: %w", err), http.StatusInternalServerError)
		return
	}
	statusCode, err := s.Login(rollno, body.Password)
	if err != nil {
		api.HttpError(w, fmt.Errorf("failed to login: %w", err), http.StatusInternalServerError)
		return
	}

	if statusCode != http.StatusOK {
		api.HttpError(w, fmt.Errorf("login failed with status code: %d", statusCode), http.StatusUnauthorized)
		return
	}

	response := AuthResponse{
		Status:  "success",
		Message: "Login successful",
	}
	
	api.WriteResponseBody(w, response)
}
