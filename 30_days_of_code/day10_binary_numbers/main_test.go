package main_test

import (
	dojo "dojo"
	"testing"
)

func TestCountNearFirstBitFromBinaryNumberInput11111ShouldReturnCountNumber5(t *testing.T) {
	expectedBinaryNumber := 5

	actualBinaryNumber := dojo.FindMaxCountBitFromBinaryNumber("11111")

	if expectedBinaryNumber != actualBinaryNumber {
		t.Errorf("expect %d but it got %d", expectedBinaryNumber, actualBinaryNumber)
	}
}

func TestCountNearFirstBitFromBinaryNumberInput1101ShouldReturnCountNumber2(t *testing.T) {
	expectedBinaryNumber := 2

	actualBinaryNumber := dojo.FindMaxCountBitFromBinaryNumber("1101")

	if expectedBinaryNumber != actualBinaryNumber {
		t.Errorf("expect %d but it got %d", expectedBinaryNumber, actualBinaryNumber)
	}
}

func TestCountNearFirstBitFromBinaryNumberInput100ShouldReturnCountNumber1(t *testing.T) {
	expectedBinaryNumber := 1

	actualBinaryNumber := dojo.FindMaxCountBitFromBinaryNumber("100")

	if expectedBinaryNumber != actualBinaryNumber {
		t.Errorf("expect %d but it got %d", expectedBinaryNumber, actualBinaryNumber)
	}
}

func TestConvertNumberFromDecimalToBinaryInput1025ShouldReturn0100(t *testing.T) {
	expectedBinaryNumber := "10000000001"

	actualBinaryNumber := dojo.ConvertNumberFromDecimalToBinary(1025)

	if expectedBinaryNumber != actualBinaryNumber {
		t.Errorf("expect %s but it got %s", expectedBinaryNumber, actualBinaryNumber)
	}
}

func TestConvertNumberFromDecimalToBinaryInput4ShouldReturn0100(t *testing.T) {
	expectedBinaryNumber := "100"

	actualBinaryNumber := dojo.ConvertNumberFromDecimalToBinary(4)

	if expectedBinaryNumber != actualBinaryNumber {
		t.Errorf("expect %s but it got %s", expectedBinaryNumber, actualBinaryNumber)
	}
}
