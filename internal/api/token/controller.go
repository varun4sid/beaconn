package token

import (
	"fmt"
	"net/http"

	"github.com/varun4sid/beaconn-backend/internal/api"
	cfg "github.com/varun4sid/beaconn-backend/internal/config"
	"github.com/varun4sid/beaconn-backend/internal/scraper"
)

func tokenHandler(w http.ResponseWriter, r *http.Request, c *cfg.Config) {
	body, err := api.ParseRequestBody[TokenRequest](r)
	if err != nil {
		api.HttpError(w, fmt.Errorf("failed to parse request body: %w", err), http.StatusBadRequest)
		return
	}

	if body.Rollno == "" || body.Password == "" {
		api.HttpError(w, fmt.Errorf("rollno and password are required"), http.StatusBadRequest)
		return
	}

	s, err := scraper.NewSession(c)
	if err != nil {
		api.HttpError(w, fmt.Errorf("failed to create scraper session: %w", err), http.StatusInternalServerError)
		return
	}
	statusCode, err := s.Login(body.Rollno, body.Password)
	if err != nil || statusCode != http.StatusOK {
		api.HttpError(w, fmt.Errorf("failed to login: %w", err), http.StatusUnauthorized)
		return
	}

	token, err := s.GetToken()
	if err != nil && token.Status == "error" {
		api.HttpError(w, fmt.Errorf("failed to get token: %w", err), http.StatusInternalServerError)
		return
	}

	if token.Status == "failed" {
		response := TokenResponse{
			Status:  "failure",
			Message: "Token unavailable at the moment",
			Token:   scraper.TokenData{},
		}
		api.WriteResponseBody(w, response)
		return
	}

	response := TokenResponse{
		Status: "success",
		Token:  token,
	}

	fmt.Printf("USER: %s - TOKEN GENERATED SUCCESSFULLY\n", body.Rollno)

	api.WriteResponseBody(w, response)
}
