package main

import "fmt"

type address struct {
	street string
	number int
	city   string
}

func main() {
	fmt.Println("Starting...")

	address := address{
		street: "A Street",
		number: 729,
		city:   "Vancouver",
	}

	fmt.Println("Address:", address)

}
