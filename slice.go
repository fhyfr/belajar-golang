package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	// create data slice
	var slice1 = months[3:7]

	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// update array it will update slice too
	months[5] = "Baru"

	fmt.Println(slice1)

	// update slice will update array too
	slice1[2] = "Juni_Again"
	fmt.Println(months)

	// create new slice
	slice2 := months[10:]
	fmt.Println(slice2)

	slice3 := append(slice2, "Muharram")
	fmt.Println(slice3)

	// create slice from beginning
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Firman"
	newSlice[1] = "Syah"

	newSlice = append(newSlice, "Fhaya")

	fmt.Println(newSlice)

	// Copy a slice
	copySlice := make([]string, 1, cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

}
