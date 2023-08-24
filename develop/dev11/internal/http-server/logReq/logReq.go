package logReq

import (
	"log"
	"net/http"
)

func LogRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %s %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	}
}
