package main_test

import (
	dojo "dojo"
	"testing"
)

func TestFactorialInput0ShouldReturn1(t *testing.T) {
	expectedNumber := int32(1)

	actualNumber := dojo.Factorial(0)

	if expectedNumber != actualNumber {
		t.Errorf("expect %d but it got %d", expectedNumber, actualNumber)
	}
}

func TestFactorialInput3ShouldReturn6(t *testing.T) {
	expectedNumber := int32(6)

	actualNumber := dojo.Factorial(3)

	if expectedNumber != actualNumber {
		t.Errorf("expect %d but it got %d", expectedNumber, actualNumber)
	}
}
