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

func TestSumOfPowerDigitInput7ShouldReturn11(t *testing.T) {
	expectedSumOfPowerDigit := 11

	actualSumOfPowerDigit := SumOfPowerDigit(7)

	assert.Equal(t, expectedSumOfPowerDigit, actualSumOfPowerDigit)
}

func TestSumOfPowerDigitInput3ShouldReturn8(t *testing.T) {
	expectedSumOfPowerDigit := 8

	actualSumOfPowerDigit := SumOfPowerDigit(3)

	assert.Equal(t, expectedSumOfPowerDigit, actualSumOfPowerDigit)
}

func TestGetPowerOfInput10ShouldReturn1024(t *testing.T) {
	expectedPowerResult := "1024"

	actualPowerResult := GetPowerOf(10)

	assert.Equal(t, expectedPowerResult, actualPowerResult)
}

func TestGetPowerOfInput2ShouldReturn2(t *testing.T) {
	expectedPowerResult := "4"

	actualPowerResult := GetPowerOf(2)

	assert.Equal(t, expectedPowerResult, actualPowerResult)
}

func TestGetPowerOfInput1ShouldReturn1(t *testing.T) {
	expectedPowerResult := "2"

	actualPowerResult := GetPowerOf(1)

	assert.Equal(t, expectedPowerResult, actualPowerResult)
}
