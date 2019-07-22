package main_test

import (
	dojo "dojo"
	"testing"
)

func TestConcatenateStringInputHackerRankAndIsTheBestPracticeShouldReturnHackerRankIsTheBestPractice(t *testing.T) {
	expectedString := "HackerRank is the best practice"

	actualString := dojo.ConcatenateString("HackerRank ", "is the best practice")

	if expectedString != actualString {
		t.Errorf("expect '%s' but it got '%s'", expectedString, actualString)
	}
}

func TestPlusFloatNumberInput4Point0And1Point5ShouldReturn5Point5(t *testing.T) {
	expectedFloatNumber := "5.5"

	actualFloatNumber := dojo.PlusFloatNumber(4.0, 1.5)

	if expectedFloatNumber != actualFloatNumber {
		t.Errorf("expect '%s' but it got '%s'", expectedFloatNumber, actualFloatNumber)
	}
}

func TestPlusNumberInput12And4ShouldReturn16(t *testing.T) {
	expectedNumber := "16"

	actualNumber := dojo.PlusNumber(12, 4)

	if expectedNumber != actualNumber {
		t.Errorf("expect '%s' but it got '%s'", expectedNumber, actualNumber)
	}
}
