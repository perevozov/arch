package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name    string
	Method  string
	Pattern string
	Handler http.Handler
}

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Routes []Route

func main() {
	log.Printf("Server started")

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":9999", router))
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	var routes = Routes{
		Route{
			"Test",
			"POST",
			"/test",
			http.HandlerFunc(TestPost),
		},
	}
	for _, route := range routes {
		var handler http.Handler
		handler = route.Handler

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}

func TestPost(w http.ResponseWriter, r *http.Request) {
	var u LoginRequest

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Could not decode request body", 400)
		return
	}

	log.Printf("request body %v", u)

	w.Write(strToByteArray("OK"))
}

func strToByteArray(str string) []byte {
	return []byte(str)
}
