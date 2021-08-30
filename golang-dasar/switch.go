package main

import "fmt"

func main() {
	// Case : Mencetak bilangan genap atau ganjil
	number := 200114140
	mod := number % 2

	switch mod {
	case 0:
		fmt.Println("Angka", number, "adalah bilangan genap")
	case 1:
		fmt.Println("Angka", number, "adalah bilangan ganjil")
	default:
		fmt.Println("Angka yang dimasukan bukan bilangan bulat")
	}

	/**
	switch dengan short statement
	Case : Menentukan kelulusan ujian
	**/
	switch score := 80; score >= 80 {
	case true:
		fmt.Println("Selamat Anda lulus")
	case false:
		fmt.Println("Anda tidak lulus! Silakan ujian ulang")
	}

	/**
	switch tanpa kondisi
	Case : Menentukan nilai ujian
	**/
	score := 61
	switch {
	case score >= 80:
		fmt.Println("Nilai : A")
	case score >= 70 && score < 80:
		fmt.Println("Nilai : B")
	case score >= 60 && score < 70:
		fmt.Println("Nilai : C")
	case score < 60 && score >= 40:
		fmt.Println("Nilai : D")
	default:
		fmt.Println("Nilai : E")
	}
}
