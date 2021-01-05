package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/health/", func(rw http.ResponseWriter, rq *http.Request) {
		rw.Write([]byte(`{"status":"OK"}`))
	})
	http.ListenAndServe("0.0.0.0:8000", nil)
}
