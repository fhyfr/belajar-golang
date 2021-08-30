package main

import "fmt"

func main() {

	name := "FIRMAN"

	if name == "firman" {
		fmt.Println("Lowercase: ", name)
	} else if name == "FIRMAN" {
		fmt.Println("Uppercase: ", name)
	} else {
		fmt.Println("Data tidak ditemukan")
	}

	// if short statement
	if length := 21; length > 20 {
		fmt.Println("Panjang lebih dari 20 meter")
	} else if length == 20 {
		fmt.Println("Panjang sama dengan 20 meter")
	} else {
		fmt.Println("Panjang kurang dari 20 meter")
	}
}
