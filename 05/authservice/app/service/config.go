package service

import (
	"os"
)

type Config struct {
	DBUser           string
	DBPasswd         string
	DBHost           string
	DBName           string
	ListenPort       string
	EmailCheckerHost string
}

var ServiceConfig Config

func init() {
	ServiceConfig.DBUser = getEnv("DB_USER", "")
	ServiceConfig.DBPasswd = getEnv("DB_PASSWD", "")
	ServiceConfig.DBHost = getEnv("DB_HOST", "localhost")
	ServiceConfig.DBName = getEnv("DB_NAME", "")
	ServiceConfig.ListenPort = getEnv("LISTEN_PORT", "80")
	ServiceConfig.EmailCheckerHost = getEnv("EMAIL_CHECKER_HOST", "emailchecker")
}

func getEnv(name string, defaultValue string) string {
	result := os.Getenv(name)
	if result == "" {
		result = defaultValue
	}
	return result
}
