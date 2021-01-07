package service

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type LoggingResponseWriter struct {
	http.ResponseWriter
	status int
	Length int
}

func NewLoggingResponseWriter(rw http.ResponseWriter) LoggingResponseWriter {
	return LoggingResponseWriter{rw, 0, 0}
}

func (w *LoggingResponseWriter) WriteHeader(status int) {
	if status < 0 {
		panic("Response status should be positive")
	}

	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *LoggingResponseWriter) Write(b []byte) (int, error) {
	if w.status == 0 {
		w.status = 200
	}
	n, err := w.ResponseWriter.Write(b)
	w.Length += n
	return n, err
}

func (w *LoggingResponseWriter) Status() int {
	return w.status
}

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		lrw := NewLoggingResponseWriter(w)
		inner.ServeHTTP(&lrw, r)

		log.Printf(
			"%s %s %d %s",
			r.Method,
			r.RequestURI,
			lrw.Status(),
			time.Since(start),
		)

		opsProcessed.
			With(prometheus.Labels{
				"method":        r.Method,
				"endpoint":      r.RequestURI,
				"http_response": strconv.Itoa(lrw.Status()),
			}).
			Inc()
	})
}
