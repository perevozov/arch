package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println("Listening")
	log.Fatal(http.ListenAndServe(":"+getEnv("LISTEN_PORT", "80"), http.HandlerFunc(hadleHTTP)))
}

func hadleHTTP(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("X-UserId")
	if userId == "" {
		log.Println("Error: no X-UserId header")
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	userName := r.Header.Get("X-UserName")
	firstName := r.Header.Get("X-UserFirstName")
	lastName := r.Header.Get("X-UserLastName")

	fmt.Fprintf(
		w,
		"User Info\nUser Id: %s\nLogin: %s\nFirst Name: %s\nLast Name: %s\n",
		userId,
		userName,
		firstName,
		lastName,
	)
	log.Printf("Serving request for user %s\n", userId)
}

func getEnv(name string, defaultValue string) string {
	result := os.Getenv(name)
	if result == "" {
		result = defaultValue
	}
	return result
}
