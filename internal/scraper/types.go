package scraper

import "net/http"

type Session struct {
	Client  *http.Client
	baseURL string
}
