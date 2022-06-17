package service


import (
	"net/http"
	"log"
)

func SetContentTypeMiddleware(h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    	file := 
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%q", fn))
        log.Println("middleware", r.URL)
        h.ServeHTTP(w, r)
    })
}