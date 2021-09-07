/*
Struct adalah sebuah template data yang digunakan untuk
menggabungkan nol atau lebih tipe data
**
Data struct disimpan dalam field,
maka struct adalah kumpulan dari field
**
Struct mirip dengan Class pada OOP
*/

package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

// Struct methods
// Method = Function
func (customer Customer) greetings(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func (customer Customer) myAge(age int) {
	fmt.Println("I am", customer.Age, "years old.", "And you are", age, ", right?")
}

func main() {
	var firman Customer
	firman.Name = "Firmansah"
	firman.Address = "Jakarta, Indonesia"
	firman.Age = 22

	firman.greetings("Helsa")
	firman.myAge(20)

	// Struct Literals 1 (prefered)
	riski := Customer{
		Name:    "Riszki Alfiansah",
		Address: "Kebumen, Jawa Tengah",
		Age:     19,
	}

	// Struct Literals 2 (rentan error pada urutan data)
	budi := Customer{"Budi Santoso", "Bekasi, Jawa Barat", 23}

	fmt.Println(firman)
	fmt.Println(riski)
	fmt.Println(budi)
}
