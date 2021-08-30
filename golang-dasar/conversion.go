package main

import "fmt"

func main() {
	var value32 = 32786
	var value64 = int64(value32)
	var value8 = int8(value64)
	var distance = value32 % 256 // because sum of value int8 is 256

	fmt.Println(value32)
	fmt.Println(value64)
	fmt.Println(value8)
	fmt.Println(distance)

	// convert byte to string
	var name = "Golang"
	var e byte = name[0]

	var eString = string(e)

	fmt.Println(eString)

}
