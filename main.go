package main

import (
	"fmt"

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

}
