package main

import (
	"log"
	"net/http"
	"os"

	"github.com/perevozov/arch/arch-emailchecker/service"
)

func main() {
	log.Printf("Server started")

	router := service.NewRouter()

	log.Fatal(http.ListenAndServe(":"+getEnv("EMAILCHECKER_LISTEN_PORT", "80"), router))
}

func getEnv(name string, defaultValue string) string {
	result := os.Getenv(name)
	if result == "" {
		result = defaultValue
	}
	return result
}
