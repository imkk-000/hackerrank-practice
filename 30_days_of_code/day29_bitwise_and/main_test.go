package main_test

import (
	. "dojo"
	"testing"
)

func TestFindMaxBitwiseAndInputN2AndK2ShouldReturnMaxBitwiseAnd0(t *testing.T) {
	expectedMaxBitwiseAnd := 0

	actualMaxBitwiseAnd := FindMaxBitwiseAnd(2, 2)

	if expectedMaxBitwiseAnd != actualMaxBitwiseAnd {
		t.Errorf("expect %d but it got %d", expectedMaxBitwiseAnd, actualMaxBitwiseAnd)
	}
}

func TestFindMaxBitwiseAndInputN8AndK5ShouldReturnMaxBitwiseAnd4(t *testing.T) {
	expectedMaxBitwiseAnd := 4

	actualMaxBitwiseAnd := FindMaxBitwiseAnd(8, 5)

	if expectedMaxBitwiseAnd != actualMaxBitwiseAnd {
		t.Errorf("expect %d but it got %d", expectedMaxBitwiseAnd, actualMaxBitwiseAnd)
	}
}

func FindMaxBitwiseAndInputN5AndK2ShouldReturnMaxBitwiseAnd1(t *testing.T) {
	expectedMaxBitwiseAnd := 1

	actualMaxBitwiseAnd := FindMaxBitwiseAnd(5, 2)

	if expectedMaxBitwiseAnd != actualMaxBitwiseAnd {
		t.Errorf("expect %d but it got %d", expectedMaxBitwiseAnd, actualMaxBitwiseAnd)
	}
}
