package practices

import (
	"fmt"
	"testing"
)

// FizzBuzz check if number if fizz or buzz
// if mod 3 == fizz
// if mod 5 == buzz
// if mod 3 and 5 == fizz buzz
func FizzBuzz(value int) {
	for i := 1; i <= value; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Fizz Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
	return
}

func TestFizzBuzz(t *testing.T) {
	FizzBuzz(100)
	FizzBuzz(200)
}
