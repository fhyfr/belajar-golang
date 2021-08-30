package main

import "fmt"

func main() {
	var name string
	var age int

	name = "Firman"
	age = 22
	var isMarried = false
	salary := 2450.6

	fmt.Println("Nama : ", name)
	fmt.Println("Umur : ", age)
	fmt.Println("Sudah Menikah? : ", isMarried)
	fmt.Println("Gaji : $", salary)

	name = "Faiz"
	age = 25
	isMarried = true
	salary = 3250.005
	fmt.Println("Nama : ", name)
	fmt.Println("Umur : ", age)
	fmt.Println("Sudah Menikah? : ", isMarried)
	fmt.Println("Gaji : $", salary)

	// Define multiple variabel

	var (
		hoby    = "Reading"
		job     = "Software Developer"
		country = "Indonesia"
	)

	print(hoby, job, country)
}
