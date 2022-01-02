package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Firman")

	if result != "HelloFirman" {
		// error 
		Fail("Result is not Hello Firman")
	}
}

// command to running unit test in all package : go test ./...