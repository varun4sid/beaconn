package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type Config struct {
	AppEnv        string
	Port          int
	AllowedDomain string
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
	}, nil
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
