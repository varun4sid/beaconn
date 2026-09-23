package scraper

import (
	"net/http"
	"net/http/cookiejar"
	"time"
	"fmt"
	
	cfg "github.com/varun4sid/beaconn-backend/internal/config"
)

func NewSession(c *cfg.Config) *Session {
	cookieJar, _ := cookiejar.New(nil)

	return &Session{
		baseURL: fmt.Sprintf("https://%s", c.AllowedDomain),
		Client: &http.Client{
			Transport: c.Transport,
			Jar:       cookieJar,
			Timeout:   15 * time.Second,
		},
	}
}
