/*
Panic function digunakan untuk menghentikan program
**
Panic function dipanggil saat terjadi error
**
Saat panic function dipanggil, program akan berhenti
tetapi defer akan tetap diekseskusi
*/

/*
Function recover kebalikan dari panic,
berfungsi untuk menangkap data panic
dan menghentikan panic
*/

package main

import "fmt"

func endApp() {
	// menangkap data panic dengan recover
	// recover dipanggil dalam defer function agar tetap dieksekusi
	message := recover()

	if message != nil {
		fmt.Println("Terjadi Error:", message)
	}
	fmt.Println("Aplikasi selesai dijalankan")
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("APLIKASI ERROR")
	}

	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp(false)
	// ketika error true function panic dipanggil dan program berhenti
	// tetapi defer tetap dijalankan
	runApp(true)
}
