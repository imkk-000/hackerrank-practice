package main_test

import (
	. "dojo"
	"testing"
)

func TestGetStringFromIndexedInputOddAndHackerShouldReturnHce(t *testing.T) {
	expectedText := "akr"

	actualText := GetStringFromIndexed(ODD, "Hacker")

	if expectedText != actualText {
		t.Errorf("expect %s but it got %s", expectedText, actualText)
	}
}

func TestGetStringOddIndexedInputEvenAndHackerShouldReturnHce(t *testing.T) {
	expectedText := "Hce"

	actualText := GetStringFromIndexed(EVEN, "Hacker")

	if expectedText != actualText {
		t.Errorf("expect %s but it got %s", expectedText, actualText)
	}
}
