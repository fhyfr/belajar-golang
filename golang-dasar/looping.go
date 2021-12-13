package main

import "fmt"

func main() {
	// Perulangan dengan for loop
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}

	/**
	Perulangan for dengan statement
	init statement : counter := 10
	condition : counter > 0
	post statement : counter++
	**/
	for counter := 10; counter > 0; counter-- {
		fmt.Println("Perulangan ke ", counter)
	}

	// looping data collection
	slice := []string{"Satu", "Dua", "Tiga", "Empat"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// For Range
	for i, value := range slice {
		fmt.Println("Index", i, "=", value)
	}

	// OR

	for _, value := range slice {
		fmt.Println(value)
	}

	// For Range data map

	person := make(map[string]string)
	person["title"] = "Belajar Golang"
	person["author"] = "Dr. Firmansah"
	person["page"] = "125 page"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
