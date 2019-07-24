package main_test

import (
	dojo "dojo"
	"testing"
)

func TestFindPrimeNumberInput1ShouldReturnFalse(t *testing.T) {
	expectedIsPrimeNumber := false

	actualIsPrimeNumber := dojo.FindPrimeNumber(1)

	if expectedIsPrimeNumber != actualIsPrimeNumber {
		t.Errorf("expect %v but it got %v", expectedIsPrimeNumber, actualIsPrimeNumber)
	}
}

func TestFindPrimeNumberInput2ShouldReturnTrue(t *testing.T) {
	expectedIsPrimeNumber := true

	actualIsPrimeNumber := dojo.FindPrimeNumber(2)

	if expectedIsPrimeNumber != actualIsPrimeNumber {
		t.Errorf("expect %v but it got %v", expectedIsPrimeNumber, actualIsPrimeNumber)
	}
}

func TestFindPrimeNumberInput4ShouldReturnFalse(t *testing.T) {
	expectedIsPrimeNumber := false

	actualIsPrimeNumber := dojo.FindPrimeNumber(4)

	if expectedIsPrimeNumber != actualIsPrimeNumber {
		t.Errorf("expect %v but it got %v", expectedIsPrimeNumber, actualIsPrimeNumber)
	}
}

func TestFindPrimeNumberInput3ShouldReturnTrue(t *testing.T) {
	expectedIsPrimeNumber := true

	actualIsPrimeNumber := dojo.FindPrimeNumber(3)

	if expectedIsPrimeNumber != actualIsPrimeNumber {
		t.Errorf("expect %v but it got %v", expectedIsPrimeNumber, actualIsPrimeNumber)
	}
}
