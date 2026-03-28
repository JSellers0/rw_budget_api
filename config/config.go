package config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	GinMode string
	ApiHost string
	ApiPort string
	DbHost  string
	DbPort  string
	DbUser  string
	DbPswd  string
	DbName  string
	LogDir  string
)

func Init() {
	_ = godotenv.Load()
	GinMode = getEnv("GIN_MODE", "debug")
	ApiHost = getEnv("API_HOST", "127.0.0.1")
	ApiPort = getEnv("API_PORT", "8081")
	DbHost = getEnv("DB_HOST", "192.168.40.101")
	DbPort = getEnv("DB_PORT", "3307")
	DbUser = getEnv("DB_USER", "svc_rw_budget")
	DbPswd = getEnv("DB_PSWD", "FAIL")
	DbName = getEnv("DB_NAME", "rw_budget_dev")
	LogDir = getEnv("LOG_DIR", "${HOME}/projects/rw_budget_api/logs/dev")
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
