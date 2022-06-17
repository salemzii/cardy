package main

import (
	"github.com/salemzii/cardy/service"
	"github.com/gorilla/mux"

	"net/http"
	"log"
)

func main() {

	
	router := mux.NewRouter().StrictSlash(false)
	router.Use(service.SetContentTypeMiddleware)
	http.Handle("/", router)

	router.HandleFunc("/cars/owner/{name}", service.GetAllCarsWithSameOwner)
	router.HandleFunc("/cars/manufacturer/{manufactor}", service.GetAllCarsWithSameManufacturer)
	router.HandleFunc("/cars/city/{city}", service.GetCarsByCity)

	log.Println("Starting http server on port :8080")
	http.ListenAndServe(":8080", router)
}
