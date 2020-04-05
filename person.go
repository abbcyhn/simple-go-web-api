package main

import (
	"time"
)

//Person ...
type Person struct {
	ID       int       `json:"Id"`
	Name     string    `json:"Name"`
	LastName string    `json:"LastName"`
	Birthday time.Time `json:"Birthday"`
	Weight   float64   `json:"Weight"`
	Height   float64   `json:"Height"`
}

//Persons ...
type Persons struct {
	Persons []Person `json:"persons"`
}
