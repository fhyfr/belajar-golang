/*
Defer function adalah function yang bisa dijadwalkan untuk dieksekusi
setelah sebuah function selesai di eksekusi
**
Defer function akan selalu dieksekusi walaupun terjadi error
di function yang dieksekusi
*/

package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function divide")
}

func divide(value int) {
	defer logging() // function logging akan selalu dieksekusi
	fmt.Println("Run Function Divide")
	result := 10 / value
	fmt.Println("Result", result)
}

func main() {
	divide(2)
	divide(0) // function logging akan dieksekusi walaupun function divide error
}
