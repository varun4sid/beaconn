package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Config struct {
	AppEnv        string
	Port          int
	AllowedOrigin string
	ClientDomain  string
	Transport     *http.Transport
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("Load .env")
	}

	portText := getEnv("PORT", "8080")
	port, err := strconv.Atoi(portText)
	if err != nil || port < 1 || port > 65535 {
		return nil, fmt.Errorf("invalid PORT %q", portText)
	}

	return &Config{
		AppEnv:        getEnv("APP_ENV", "no_env"),
		Port:          port,
		AllowedOrigin: getEnv("ALLOWED_ORIGIN", ""),
		ClientDomain:  getEnv("CLIENT_DOMAIN", ""),
		Transport:     newTransport(),
	}, nil
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func newTransport() *http.Transport {
	tr := http.DefaultTransport.(*http.Transport).Clone()

	tr.MaxIdleConns = 100
	tr.MaxIdleConnsPerHost = 20
	tr.MaxConnsPerHost = 40
	tr.MaxIdleConns = 100
	tr.MaxIdleConnsPerHost = 20
	tr.MaxConnsPerHost = 40

	tr.IdleConnTimeout = 60 * time.Second
	tr.IdleConnTimeout = 60 * time.Second

	tr.TLSHandshakeTimeout = 10 * time.Second
	tr.ResponseHeaderTimeout = 10 * time.Second
	tr.ExpectContinueTimeout = 1 * time.Second
	tr.TLSHandshakeTimeout = 10 * time.Second
	tr.ResponseHeaderTimeout = 10 * time.Second
	tr.ExpectContinueTimeout = 1 * time.Second

	tr.DisableKeepAlives = false
	tr.DisableKeepAlives = false

	return tr
}
