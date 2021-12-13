/*
Type assertions digunakan untuk mengubah tipe data
yang dikembalikan oleh interface kosong
*/

package main

import "fmt"

func random() interface{} {
	return "OK"
}

func main() {
	// var result interface{} = random()
	// var resultString string = result.(string)
	// fmt.Println(resultString)

	result := random()

	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	default:
		fmt.Println("Unknown data type")
	}
}
