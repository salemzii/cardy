package service


import (
	"net/http"
	"log"
)

func SetContentTypeMiddleware(h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
        log.Println("middleware", r.URL)
        h.ServeHTTP(w, r)
    })
}