package main

import (
	"errors"
	"fmt"
)

func Divide(pembilang int, penyebut int) (int, error) {
	if penyebut == 0 {
		return 0, errors.New("pembagian dengan NOL")
	} else {
		return pembilang / penyebut, nil
	}
}

func main() {
	var exError error = errors.New("Ups error")
	fmt.Println(exError.Error())

	result, err := Divide(100, 0)

	if err == nil {
		fmt.Println("Hasil", result)
	} else {
		fmt.Println("Error", err.Error())
	}
}
