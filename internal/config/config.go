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
	AllowedDomain string
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
		AppEnv:        getEnv("APP_ENV", "development"),
		Port:          port,
		AllowedDomain: getEnv("ALLOWED_DOMAIN", ""),
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
	return &http.Transport{
		MaxIdleConns:        50,
		MaxIdleConnsPerHost: 50,
		MaxConnsPerHost:     40,
		IdleConnTimeout:     180 * time.Second,
	}
}
