package main

import "fmt"

func main() {
	type noKTP string
	type married bool

	var noKtpfirman noKTP = "33301222821121"
	var isMarried married = false
	fmt.Println(noKtpfirman)
	fmt.Println("Sudah Menikah?", isMarried)
}
