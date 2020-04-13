package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", Index)

	router.HandleFunc("/GetById/{id}", GetByID).Methods("GET")
	router.HandleFunc("/GetByWeight/{weight}", GetByWeight).Methods("GET")
	router.HandleFunc("/GetByHeight/{height}", GetByHeight).Methods("GET")

	router.HandleFunc("/AddPerson", AddPerson).Methods("POST")
	router.HandleFunc("/EditPerson", EditPerson).Methods("POST")
	router.HandleFunc("/RemovePerson/{id}", RemovePerson).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
