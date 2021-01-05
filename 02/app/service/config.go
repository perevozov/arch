package service

import (
	"os"
)

type Config struct {
	DBUser     string
	DBPasswd   string
	DBHost     string
	DBName     string
	ListenPort string
}

var ServiceConfig Config

func init() {
	ServiceConfig.DBUser = getEnv("HW02_DB_USER", "")
	ServiceConfig.DBPasswd = getEnv("HW02_DB_PASSWD", "")
	ServiceConfig.DBHost = getEnv("HW02_DB_HOST", "localhost")
	ServiceConfig.DBName = getEnv("HW02_DB_NAME", "")
	ServiceConfig.ListenPort = getEnv("HW02_LISTEN_PORT", "80")
}

func getEnv(name string, defaultValue string) string {
	result := os.Getenv(name)
	if result == "" {
		result = defaultValue
	}
	return result
}
