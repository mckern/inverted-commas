package main

import "testing"

func TestDoubleSmartQuotes(t *testing.T) {
	doubleQuoteStr := "\u201c, \u201d, \u201e, \u201f, \u2033, \u2036"
	actualDoubleResult := doubleQuotes.ReplaceAllString(doubleQuoteStr, "x")
	expectedDoubleResult := "x, x, x, x, x, x"

	if actualDoubleResult != expectedDoubleResult {
		t.Fatalf("Expected \"%s\" but got \"%s\"", expectedDoubleResult, actualDoubleResult)
	}
}
func TestSingleSmartQuotes(t *testing.T) {
	singleQuoteStr := "\u2018, \u2019, \u201a, \u201b, \u2032, \u2035"
	actualSingleResult := singleQuotes.ReplaceAllString(singleQuoteStr, "x")
	expectedSingleResult := "x, x, x, x, x, x"

	if actualSingleResult != expectedSingleResult {
		t.Fatalf("Expected \"%s\" but got \"%s\"", expectedSingleResult, actualSingleResult)
	}
}
