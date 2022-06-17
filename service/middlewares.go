package service


import (
	"net/http"
	"log"
	"io"
	"bytes"
	"io/ioutil"
	"fmt"
)

func SetContentTypeMiddleware(h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    	data, err := ioutil.ReadFile("/home/salem/Desktop/cardy/data.csv")
		if err != nil {
			fmt.Println(err)
		}

		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=data.csv")
		
		io.Copy(w, bytes.NewReader(data))
        log.Println("middleware", r.URL)
        h.ServeHTTP(w, r)
    })
}