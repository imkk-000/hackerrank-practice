package main_test

import (
	"bytes"
	. "dojo"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainInputFile00ShouldBeOutputFile00(t *testing.T) {
	expectedResult, _ := ioutil.ReadFile("output/output00.txt")
	expectedResult = bytes.TrimSpace(expectedResult)
	cmd := exec.Command("go", "run", "main.go")
	cmd.Stdin, _ = os.Open("input/input00.txt")

	actualResult, _ := cmd.Output()
	actualResult = bytes.TrimSpace(actualResult)

	assert.Equal(t, expectedResult, actualResult)
}

func TestFindHiglyDivisibleTriangularNumberInput4ShouldBe28(t *testing.T) {
	expectedNumber := 28

	actualNumber := FindHiglyDivisibleTriangularNumber(4)

	assert.Equal(t, expectedNumber, actualNumber)
}

func TestFindHiglyDivisibleTriangularNumberInput3ShouldBe6(t *testing.T) {
	expectedNumber := 6

	actualNumber := FindHiglyDivisibleTriangularNumber(3)

	assert.Equal(t, expectedNumber, actualNumber)
}

func TestFindHiglyDivisibleTriangularNumberInput2ShouldBe6(t *testing.T) {
	expectedNumber := 6

	actualNumber := FindHiglyDivisibleTriangularNumber(2)

	assert.Equal(t, expectedNumber, actualNumber)
}

func TestFindHiglyDivisibleTriangularNumberInput1ShouldBe3(t *testing.T) {
	expectedNumber := 3

	actualNumber := FindHiglyDivisibleTriangularNumber(1)

	assert.Equal(t, expectedNumber, actualNumber)
}

func TestSumOfInput1ShouldBe1(t *testing.T) {
	expectedSum := 1

	actualSum := SumOf(1)

	assert.Equal(t, expectedSum, actualSum)
}

func TestSumOfInput10ShouldBe55(t *testing.T) {
	expectedSum := 55

	actualSum := SumOf(10)

	assert.Equal(t, expectedSum, actualSum)
}

func TestFindFactorOfInput36ShouldReturnFators(t *testing.T) {
	expectedFators := []int{1, 36, 2, 18, 3, 12, 4, 9, 6}

	actualFactors := FindFactorOf(36)

	assert.Equal(t, expectedFators, actualFactors)
}

func TestFindFactorOfInput6ShouldReturnFators(t *testing.T) {
	expectedFators := []int{1, 6, 2, 3}

	actualFactors := FindFactorOf(6)

	assert.Equal(t, expectedFators, actualFactors)
}

func TestFindFactorOfInput3ShouldReturnFators(t *testing.T) {
	expectedFators := []int{1, 3}

	actualFactors := FindFactorOf(3)

	assert.Equal(t, expectedFators, actualFactors)
}

func TestFindFactorOfInput1ShouldReturnFators(t *testing.T) {
	expectedFators := []int{1}

	actualFactors := FindFactorOf(1)

	assert.Equal(t, expectedFators, actualFactors)
}
