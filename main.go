package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", Index)

	router.HandleFunc("/GetById/{id}", GetByID)
	router.HandleFunc("/GetByWeight/{weight}", GetByWeight)
	router.HandleFunc("/GetByHeight/{height}", GetByHeight)

	router.HandleFunc("/AddPerson", AddPerson)
	router.HandleFunc("/EditPerson", EditPerson)
	router.HandleFunc("/RemovePerson/{id}", RemovePerson)

	log.Fatal(http.ListenAndServe(":8080", router))
}
