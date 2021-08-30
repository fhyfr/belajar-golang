package main

import "fmt"

func main() {
	// continue : skip nilai tertentu
	for i := 0; i < 12; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Println("Bilangan Ganjil :", i)
	}

	// break : menghentikan perulangan
	for x := 0; x < 100; x++ {
		if x == 50 {
			break
		}

		fmt.Println("Perulangan ke ", x)
	}
}
