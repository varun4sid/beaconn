package token

import (
	"github.com/varun4sid/beaconn-backend/internal/scraper"
)

type TokenRequest struct {
	Rollno   string `json:"rollno"`
	Password string `json:"password"`
}

type TokenResponse struct {
	Status  string            `json:"status"`
	Message string            `json:"message"`
	Token   scraper.TokenData `json:"token"`
}
