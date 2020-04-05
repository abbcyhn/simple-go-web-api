package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

//SelectPersons ...
func SelectPersons() Persons {
	var allPersons Persons

	jsonFile, _ := os.Open("persons.json")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &allPersons)

	return allPersons
}

//SelectByID ...
func SelectByID(id int) Person {
	var selectedPerson Person

	var allPersons = SelectPersons()

	for i := 0; i < len(allPersons.Persons); i++ {
		if allPersons.Persons[i].ID == id {
			selectedPerson = allPersons.Persons[i]
			break
		}
	}

	return selectedPerson
}

//SelectByWeight ...
func SelectByWeight(weight float64) []Person {
	var selectedPersons []Person

	allPersons := SelectPersons()

	for i := 0; i < len(allPersons.Persons); i++ {
		if allPersons.Persons[i].Weight == weight {
			selectedPersons = append(selectedPersons, allPersons.Persons[i])
		}
	}

	return selectedPersons
}

//SelectByHeight ...
func SelectByHeight(height float64) []Person {
	var selectedPersons []Person

	allPersons := SelectPersons()

	for i := 0; i < len(allPersons.Persons); i++ {
		if allPersons.Persons[i].Height == height {
			selectedPersons = append(selectedPersons, allPersons.Persons[i])
		}
	}

	return selectedPersons
}

//InsertPerson ...
func InsertPerson(person Person) {
	var allPersons = SelectPersons()

	var length = len(allPersons.Persons)

	person.ID = length + 1

	allPersons.Persons = append(allPersons.Persons, person)

	file, _ := json.MarshalIndent(allPersons, "", " ")

	_ = ioutil.WriteFile("persons.json", file, 0644)
}

//UpdateByID ...
func UpdateByID(person Person) {
	allPersons := SelectPersons()

	index := person.ID - 1

	allPersons.Persons[index] = person

	file, _ := json.MarshalIndent(allPersons, "", " ")

	_ = ioutil.WriteFile("persons.json", file, 0644)
}

//DeleteByID ...
func DeleteByID(id int) {
	allPersons := SelectPersons()

	index := id - 1

	length := len(allPersons.Persons)

	//remove from slice
	allPersons.Persons[index] = allPersons.Persons[length-1]
	allPersons.Persons[length-1] = Person{}
	allPersons.Persons = allPersons.Persons[:length-1]

	file, _ := json.MarshalIndent(allPersons, "", " ")

	_ = ioutil.WriteFile("persons.json", file, 0644)
}
