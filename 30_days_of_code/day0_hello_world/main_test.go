package main_test

import (
	main "dojo"
	"testing"
)

func TestShowHelloInputHello30DaysOfCodeShouldReturnHelloWorldWithHello30DaysOfCode(t *testing.T) {
	expectedResult := "Hello, World.\nWelcome to 30 Days of Code!"

	actualResult := main.ShowHello("Welcome to 30 Days of Code!")

	if expectedResult != actualResult {
		t.Errorf("expect '%s' but it got '%s'", expectedResult, actualResult)
	}
}
