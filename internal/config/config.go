package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv            string
	Port              string
	AllowedOrigin     string
	HttpClientTimeout time.Duration
}

func Load() *Config {
	_ = godotenv.Load()

	timeoutStr := getEnv("HTTP_CLIENT_TIMEOUT", "10s")
	timeout, err := time.ParseDuration(timeoutStr)
	if err != nil {
		log.Fatalf("Invalid HTTP_CLIENT_TIMEOUT: %v", err)
	}

	return &Config{
		AppEnv: getEnv("APP_ENV", "development"),
		Port: getEnv("PORT", "8080"),
		AllowedOrigin: getEnv("ALLOWED_ORIGIN", "*"),
		HttpClientTimeout: timeout,
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}