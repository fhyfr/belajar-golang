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

// function return multiple value
func getBiodata() (string, string, int) {
	return "Firman", "Sah", 22
}

// named return value
func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Firman"
	middleName = "Syah"
	lastName = "Fhaya"

	return
}

func main() {
	for i := 0; i < 3; i++ {
		sayHello()
	}

	// memanggil function dengan argument
	sayHelloTo("Firman", "sah")
	fmt.Println(getNumber(100))

	firstName, _, _ := getBiodata()

	fmt.Println(firstName)
	fmt.Println(getCompleteName())
}
