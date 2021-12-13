package helper

import "fmt"

// access modifier
// diawali huruf kecil tidak bisa diakses dari luar package
// diawali huruf besar bisa diakses dari luar package

var version = "1.0.0"
var Application = "golang"

func SayHello(name string) {
	fmt.Println("Hello", name)
}

func fullName(firstName string, lastName string) {
	fmt.Println(firstName, lastName)
}