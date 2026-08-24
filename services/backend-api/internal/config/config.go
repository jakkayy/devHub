package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	// Server Configuration
	Port        string
	Environment string

	// Database Configuration
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string

	// Redis Configuration
	RedisHost     string
	RedisPassword string

	// External Integrations Credentials
	GitHubPAT         string
	DiscordWebhookURL string
	GoogleSheetID     string
	MiroOAuthToken    string

	// Application Timeouts & Limits
	APITimeoutSeconds int
}

func LoadConfig() *Config {
	return &Config{
		Port:              getEnv("PORT", "8080"),
		Environment:       getEnv("ENV", "development"),
		DBHost:            getEnv("DB_HOST", "localhost"),
		DBPort:            getEnv("DB_PORT", "5433"),
		DBUser:            getEnv("DB_USER", "devhub"),
		DBPassword:        getEnv("DB_PASSWORD", "devhubpass"),
		DBName:            getEnv("DB_NAME", "devhub_db"),
		DBSSLMode:         getEnv("DB_SSLMODE", "disable"),
		RedisHost:         getEnv("REDIS_HOST", "localhost:6380"),
		RedisPassword:     getEnv("REDIS_PASSWORD", ""),
		GitHubPAT:         getEnv("GITHUB_PAT", ""),
		DiscordWebhookURL: getEnv("DISCORD_WEBHOOK_URL", ""),
		GoogleSheetID:     getEnv("GOOGLE_SHEET_ID", ""),
		MiroOAuthToken:    getEnv("MIRO_OAUTH_TOKEN", ""),
		APITimeoutSeconds: getEnvAsInt("API_TIMEOUT_SECONDS", 10),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		return value
	}
	log.Printf("[Config] Key '%s' not set, using default: '%s'", key, fallback)
	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	valueStr := getEnv(key, "")
	if valueStr == "" {
		return fallback
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Printf("[Config] Invalid integer value for '%s': %v, using default: %d", key, err, fallback)
		return fallback
	}
	return value
}
