package main_test

import (
	dojo "dojo"
	"testing"
)

func TestFindMaxSumOfHourglassInputArrayShouldReturnMaxNumber0(t *testing.T) {
	expectedMaxNumber := int32(0)
	input2DArray := [][]int32{
		{0, 0, 0, 0, 0, 0},
		{0, 9, 9, 9, 0, 0},
		{0, 9, 9, 9, 0, 0},
		{0, 9, 9, 9, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
	}

	actualMaxNumber := dojo.FindMaxSumOfHourglass(input2DArray)

	if expectedMaxNumber != actualMaxNumber {
		t.Errorf("expect %d but it got %d", expectedMaxNumber, actualMaxNumber)
	}
}

func TestFindMaxSumOfHourglassInputArrayShouldReturnMaxNumber19(t *testing.T) {
	expectedMaxNumber := int32(19)
	input2DArray := [][]int32{
		{1, 1, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 0},
		{0, 0, 2, 4, 4, 0},
		{0, 0, 0, 2, 0, 0},
		{0, 0, 1, 2, 4, 0},
	}

	actualMaxNumber := dojo.FindMaxSumOfHourglass(input2DArray)

	if expectedMaxNumber != actualMaxNumber {
		t.Errorf("expect %d but it got %d", expectedMaxNumber, actualMaxNumber)
	}
}
