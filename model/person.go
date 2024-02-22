package model

import "time"

type Person struct {
	Name     string
	Address  Address
	BirthDay time.Time
	Age      int
}

func (person *Person) MethodPersonCalculate() {
	year := person.BirthDay.Year()
	currently := time.Now().Year()
	person.Age = currently - year
}

func CalculateAge(person Person) int {
	year := person.BirthDay.Year()
	currentlyYear := time.Now().Year()
	return currentlyYear - year
}
