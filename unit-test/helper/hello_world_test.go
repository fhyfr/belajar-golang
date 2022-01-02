package helper

import (
	"testing"
	"fmt"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Firman")

	if result != "Hello Firman" {
		// error from unit test
		t.Error("Result must be 'Hello Firman'")
	}

	fmt.Println("Testing Done")
}

// command to running unit test in all package : go test ./...

/*

Failed Error: 
- t.Fail()
- t.FailNow()
- t.Error(args...)
- t.Fatal(args...)

*/