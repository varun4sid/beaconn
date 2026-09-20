package config

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type Config struct {
	AppEnv string
	Port   uint64
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	port, err := strconv.ParseUint(getEnv("PORT", "8080"), 10, 0)
	if err != nil {
		port = 8080
	}

	return &Config{
		AppEnv: getEnv("APP_ENV", "development"),
		Port:   port,
	}, nil
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}