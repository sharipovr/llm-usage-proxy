package config

import (
	"os"
)

// Config is the subset of environment variables the server needs so far.
// We will it on every step; keeping it small for now makes the difference obvious.
type Config struct {
	ListenAddr string
	LogLevel   string
}

// Load reads configuration from the environment with sensible fallbacks.
func Load() *Config {
	return &Config{
		ListenAddr: getEnv("LISTEN_ADDR", ":8080"),
		LogLevel:   getEnv("LOG_LEVEL", "info"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
