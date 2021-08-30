package main

import "fmt"

func main() {
	// create map data
	person := map[string]string{
		"name":    "Firman",
		"address": "Jakarta",
		"age":     "21",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	// update map values
	person["name"] = "Helsa"

	fmt.Println(person)

	// create new map
	book := make(map[string]string)
	book["title"] = "Belajar Pemrograman Golang Dasar"
	book["author"] = "Firmansah"
	book["page"] = "128"
	book["price"] = "$20"

	fmt.Println(book)

	// delete map data
	delete(book, "price")

	fmt.Println(book)
	fmt.Println(len(book))
}
