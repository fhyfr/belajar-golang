/*
Interface adalah tipe data abstract,
tidak memiliki implementasi langsung
**
Interface digunakan sebagai kontrak
*/

package main

import "fmt"

type HasName interface {
	GetName() string
}

/*
Interface kosong
**
Berfungsi seperti superclass pada OOP
***
Contoh: fungsi Println()
*/

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

func greetings(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	var firman Person
	firman.Name = "Firman"

	greetings(firman)

	cat := Animal{
		Name: "Mpus",
	}

	greetings(cat)

	// interface kosong
	var data interface{} = Ups(2)
	fmt.Println(data)
}
