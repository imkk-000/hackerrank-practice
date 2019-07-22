package main_test

import (
	dojo "dojo"
	"testing"
)

func TestShowWeirdNumberShouldReturnWeird(t *testing.T) {
	expectedResult := "Weird"
	inputNumbers := []int32{1, 3, 5, 6, 8, 20}

	for _, inputNumber := range inputNumbers {
		actualResult := dojo.ShowWeirdNumber(inputNumber)

		if expectedResult != actualResult {
			t.Errorf("expect %d is weird but it got %d is %s",
				inputNumber, inputNumber, actualResult)
		}
	}
}

func TestShowWeirdNumberShouldReturnNotWeird(t *testing.T) {
	expectedResult := "Not Weird"
	inputNumbers := []int32{2, 4, 22, 24}

	for _, inputNumber := range inputNumbers {
		actualResult := dojo.ShowWeirdNumber(inputNumber)

		if expectedResult != actualResult {
			t.Errorf("expect %d is weird but it got %d is %s",
				inputNumber, inputNumber, actualResult)
		}
	}
}
