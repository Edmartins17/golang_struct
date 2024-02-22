package main

import (
	"fmt"
	"time"

	"github.com/Edmartins17/golang_struct/model"
)

func main() {
	fmt.Println("Starting...")

	address := model.Address{
		Street: "A Street",
		Number: 729,
		City:   "Vancouver",
	}

	fmt.Println("Address:", address)

	address.Number = 10
	fmt.Println("Number changing:", address.Number)

	Person := model.Person{
		Name:     "Ed",
		Address:  address,
		BirthDay: time.Date(1977, 06, 17, 0, 0, 0, 0, time.Local),
	}

	fmt.Println(Person)

	age := model.CalculateAge(Person)
	fmt.Println(age)

	Person.MethodPersonCalculate()
	fmt.Println(age)

	vehicle := model.Vehicle{
		Year:  2022,
		Plate: "AAA-2323",
		Model: "Yamaha",
	}

	motorCycle := model.MotorCycle{
		Vehicle:          vehicle,
		CylinderCapacity: 1000,
	}

	fmt.Println(motorCycle)

}
