package main_test

import (
	"bytes"
	. "dojo"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecInputFile00ShouldBeOutputFile00(t *testing.T) {
	expectedResult, _ := ioutil.ReadFile("output/output00.txt")
	expectedResult = bytes.TrimSpace(expectedResult)
	os.Stdin, _ = os.Open("input/input00.txt")
	reader, writer, _ := os.Pipe()
	os.Stdout = writer

	Exec()
	writer.Close()
	defer reader.Close()
	actualResult, _ := ioutil.ReadAll(reader)
	actualResult = bytes.TrimSpace(actualResult)

	assert.Equal(t, expectedResult, actualResult)
}

func TestFindPrimeFromIndexOfInput6ShouldReturn13(t *testing.T) {
	expectedPrimeNumber := 13

	actualPrimeNumber := FindPrimeNumberFromIndexOf(6)

	assert.Equal(t, expectedPrimeNumber, actualPrimeNumber)
}

func TestFindPrimeFromIndexOfInput3ShouldReturn5(t *testing.T) {
	expectedPrimeNumber := 5

	actualPrimeNumber := FindPrimeNumberFromIndexOf(3)

	assert.Equal(t, expectedPrimeNumber, actualPrimeNumber)
}

func TestFindPrimeFromIndexOfInput2ShouldReturn3(t *testing.T) {
	expectedPrimeNumber := 3

	actualPrimeNumber := FindPrimeNumberFromIndexOf(2)

	assert.Equal(t, expectedPrimeNumber, actualPrimeNumber)
}

func TestFindPrimeFromIndexOfInput1ShouldReturn2(t *testing.T) {
	expectedPrimeNumber := 2

	actualPrimeNumber := FindPrimeNumberFromIndexOf(1)

	assert.Equal(t, expectedPrimeNumber, actualPrimeNumber)
}

func TestIsPrimeInput4ShouldReturnFalse(t *testing.T) {
	expectedIsPrime := false

	actaulIsPrime := IsPrime(4)

	assert.Equal(t, expectedIsPrime, actaulIsPrime)
}

func TestIsPrimeInput3ShouldReturnTrue(t *testing.T) {
	expectedIsPrime := true

	actaulIsPrime := IsPrime(3)

	assert.Equal(t, expectedIsPrime, actaulIsPrime)
}

func TestIsPrimeInput2ShouldReturnTrue(t *testing.T) {
	expectedIsPrime := true

	actaulIsPrime := IsPrime(2)

	assert.Equal(t, expectedIsPrime, actaulIsPrime)
}

func TestIsPrimeInput1ShouldReturnFalse(t *testing.T) {
	expectedIsPrime := false

	actaulIsPrime := IsPrime(1)

	assert.Equal(t, expectedIsPrime, actaulIsPrime)
}
