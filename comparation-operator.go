package main

import "fmt"

func main() {
	var name1 = "Firman"
	var name2 = "firman"

	var isSame = name1 == name2
	var isDifferent = name1 != name2

	fmt.Println(isSame)
	fmt.Println(isDifferent)
}
