package model

import "time"

type Person struct {
	Name     string
	Address  Address
	BirthDay time.Time
}

func CalculateAge(person Person) int {
	year := person.BirthDay.Year()
	currentlyYear := time.Now().Year()
	return currentlyYear - year
}
