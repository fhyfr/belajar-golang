package main

import "fmt"

func sayHello() {
	fmt.Println("Halo Dunia")
}

// function dengan parameter
func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

// function return value
func getNumber(number int) int {
	return number
}

func main() {
	for i := 0; i < 3; i++ {
		sayHello()
	}

	// memanggil function dengan argument
	sayHelloTo("Firman", "sah")
	fmt.Println(getNumber(100))
}
