package middlewares

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time := time.Now()
		next.ServeHTTP(w, r)
		log.Println(r.Method, r.URL.Path, time)
	})
}
