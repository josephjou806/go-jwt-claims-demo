package config

import (
	"os"
	"strings"
)

type Config struct {
	Port              string
	JWTSecret         string
	CORSAllowedOrigin []string
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func Load() Config {
	corsStr := getEnv("CORS_ALLOWED_ORIGINS", "*")
	var origins []string
	for _, o := range strings.Split(corsStr, ",") {
		origins = append(origins, strings.TrimSpace(o))
	}

	return Config{
		Port:              getEnv("PORT", "8080"),
		JWTSecret:         getEnv("JWT_SECRET", "dev-secret-change-me"),
		CORSAllowedOrigin: origins,
	}
}
