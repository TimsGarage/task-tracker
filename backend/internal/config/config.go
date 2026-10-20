package config

import "os"

type Config struct {
	Port          string
	DatabasePath  string
	AllowedOrigin string
}

func Load() Config {
	return Config{
		Port:          getEnv("PORT", "8080"),
		DatabasePath:  getEnv("DATABASE_PATH", "todos.db"),
		AllowedOrigin: getEnv("ALLOWED_ORIGIN", "http://localhost:5173"),
	}
}

func (config Config) Address() string {
	return ":" + config.Port
}

func getEnv(name string, fallback string) string {
	value := os.Getenv(name)
	if value == "" {
		return fallback
	}

	return value
}
