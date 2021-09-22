package main

import "fmt"

// membuat pointer pada method
type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	firman := Man{"Firman"}

	firman.Married()

	fmt.Println(firman.Name)
}
