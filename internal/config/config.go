package config

import "os"

type Config struct {
	AppPort string
	DatabaseURL string
}

func Load() *Config {
	return &Config{
		AppPort:    getEnv("APP_PORT", "8080"),
		DatabaseURL: getEnv("DATABASE_URL", "postgres://postgres:password@localhost:5432/postgres?sslmode=disable"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}

	return fallback
}