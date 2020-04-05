package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Index ...
func Index(w http.ResponseWriter, r *http.Request) {
	var allPersons = SelectPersons()

	fmt.Fprintln(w, "Length of persons: ", len(allPersons.Persons))
}

//GetByID ...
func GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	idParam := vars["id"]

	id, _ := strconv.Atoi(idParam)

	selectedPerson := SelectByID(id)

	json.NewEncoder(w).Encode(selectedPerson)
}

//GetByWeight ...
func GetByWeight(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	weightParam := vars["weight"]

	weight, _ := strconv.ParseFloat(weightParam, 32)

	selectedPersons := SelectByWeight(weight)

	json.NewEncoder(w).Encode(selectedPersons)
}

//GetByHeight ...
func GetByHeight(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	heightParam := vars["height"]

	height, _ := strconv.ParseFloat(heightParam, 32)

	selectedPersons := SelectByHeight(height)

	json.NewEncoder(w).Encode(selectedPersons)
}

//AddPerson ...
func AddPerson(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	person := Person{}

	json.Unmarshal(body, &person)

	InsertPerson(person)

	fmt.Fprintln(w, "Person added")
}

//EditPerson ...
func EditPerson(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	person := Person{}

	json.Unmarshal(body, &person)

	UpdateByID(person)

	fmt.Fprintln(w, "Person edited")
}

//RemovePerson ...
func RemovePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	idParam := vars["id"]

	id, _ := strconv.Atoi(idParam)

	DeleteByID(id)

	fmt.Fprintln(w, "Person removed")
}
