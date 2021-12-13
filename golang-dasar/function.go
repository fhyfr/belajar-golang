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

// function variadic
func sumAll(numbers ...int) int {
	total := 0

	for _, value := range numbers {
		total += value
	}

	return total
}

// function value
func getGoodBye(name string) string {
	return "Good Bye " + name
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

	total := sumAll()
	total = sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(total)

	// menggunakan slice sebagai variadic function
	slice := []int{10, 20, 30, 40, 50, 60}
	total = sumAll(slice...)
	fmt.Println(total)

	// menggunakan function value
	sayGoodBye := getGoodBye("Firman")
	fmt.Println(sayGoodBye)
}
