package main

import (
	"bytes"
	"testing"
	"fmt"
)

func TestValidateInputToReturnTrue(t *testing.T) {
	str := "word up!"
	buffer := new(bytes.Buffer)
	fmt.Fprint(buffer, str)

	newReader, _ := validateInput(buffer)

	actualBuffer := new(bytes.Buffer)
	actualBuffer.ReadFrom(newReader)
	actualResult := actualBuffer.String()

	if actualResult != str {
		t.Fatalf("Expected \"%s\" but got \"%s\"", str, actualResult)
	}
}
