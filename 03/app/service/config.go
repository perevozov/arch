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
	ServiceConfig.DBUser = getEnv("HW03_DB_USER", "")
	ServiceConfig.DBPasswd = getEnv("HW03_DB_PASSWD", "")
	ServiceConfig.DBHost = getEnv("HW03_DB_HOST", "localhost")
	ServiceConfig.DBName = getEnv("HW03_DB_NAME", "")
	ServiceConfig.ListenPort = getEnv("HW03_LISTEN_PORT", "80")
	ServiceConfig.EmailCheckerHost = getEnv("EMAIL_CHECKER_HOST", "emailchecker")
}

func getEnv(name string, defaultValue string) string {
	result := os.Getenv(name)
	if result == "" {
		result = defaultValue
	}
	return result
}
