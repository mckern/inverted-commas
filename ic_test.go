package main

import "testing"

func TestDoubleSmartQuotes(t *testing.T) {
	doubleQuoteStr := "\u201c, \u201d, \u201e, \u201f, \u2033, \u2036"
	doubleQuoteLiterals := `“, ”, „, ‟, ″, ‶`

	actualDoubleResult := doubleQuotes.ReplaceAllString(doubleQuoteStr, "x")
	actualLiteralResult := doubleQuotes.ReplaceAllString(doubleQuoteLiterals, "x")
	expectedDoubleResult := "x, x, x, x, x, x"

	if (actualDoubleResult != actualLiteralResult) ||
		(actualDoubleResult != expectedDoubleResult) {
		t.Fatalf("Expected \"%s\" but got \"%s\"", expectedDoubleResult, actualDoubleResult)
	}
}
func TestSingleSmartQuotes(t *testing.T) {
	singleQuoteStr := "\u2018 \u2019 \u201a \u201b \u2032 \u2035"
	singleQuoteLiterals := `‘ ’ ‚ ‛ ′ ‵`

	actualSingleResult := singleQuotes.ReplaceAllString(singleQuoteStr, "x")
	actualLiteralResult := singleQuotes.ReplaceAllString(singleQuoteLiterals, "x")
	expectedSingleResult := "x x x x x x"

	if (actualSingleResult != actualLiteralResult) ||
		(actualSingleResult != expectedSingleResult) {
		t.Fatalf("Expected \"%s\" but got \"%s\"", expectedSingleResult, actualSingleResult)
	}
}
