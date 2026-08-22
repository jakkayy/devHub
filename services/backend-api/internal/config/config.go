package config

import (
	"log"
	"os"
)

type Config struct {
	Port        string
	Environment string
	DBHost      string
	DBPort      string
	DBUser      string
	DBPassword  string
	DBName      string
	DBSSLMode   string
	RedisHost   string
}

func LoadConfig() *Config {
	return &Config{
		Port:        getEnv("PORT", "8080"),
		Environment: getEnv("ENV", "development"),
		DBHost:      getEnv("DB_HOST", "localhost"),
		DBPort:      getEnv("DB_PORT", "5433"),
		DBUser:      getEnv("DB_USER", "devhub"),
		DBPassword:  getEnv("DB_PASSWORD", "devhubpass"),
		DBName:      getEnv("DB_NAME", "devhub_db"),
		DBSSLMode:   getEnv("DB_SSLMODE", "disable"),
		RedisHost:   getEnv("REDIS_HOST", "localhost:6379"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		return value
	}
	log.Printf("[Config] Key %s not set, using default: %s", key, fallback)
	return fallback
}
