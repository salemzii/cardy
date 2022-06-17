package service

import (
	"github.com/salemzii/cardy/cars"
	"github.com/gorilla/mux"

	"net/http"
	"encoding/json"
	//"fmt"
	"io/ioutil"
	"log"
)

var dataStore []cars.Info

func init() {
	data, err := ioutil.ReadFile("./data.json")
	if err != nil {
		log.Println(err)
	}

	json.Unmarshal([]byte(data), &dataStore)
}

func HandlePostData(w http.ResponseWriter, r *http.Request){


}


func GetAllCarsWithSameOwner(w http.ResponseWriter, r *http.Request) {

	param := mux.Vars(r)
	name := param["name"]

	ls := cars.CarsWithOneOwner(dataStore, name)
	WriteFile(ls)
	data, err := json.Marshal(ls)
	if err != nil {
		log.Println(err)
	}

	w.Write(data)
}

func GetAllCarsWithSameManufacturer(w http.ResponseWriter, r *http.Request) {

	param := mux.Vars(r)
	manufacturer := param["manufactor"]

	ls := cars.CarWithSameManufacturer(dataStore, manufacturer)
	WriteFile(ls)

	data, err := json.Marshal(ls)
	if err != nil {
		log.Println(err)
	}

	w.Write(data)
}

func GetCarsByCity(w http.ResponseWriter, r *http.Request) {

	param := mux.Vars(r)
	city := param["city"]

	ls := cars.SearchByCity(dataStore, city)
	WriteFile(ls)
	
	data, err := json.Marshal(ls)
	if err != nil {
		log.Println(err)
	}

	w.Write(data)
}
