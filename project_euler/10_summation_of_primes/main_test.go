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
	reader, writer, _ := os.Pipe()
	os.Stdin, _ = os.Open("input/input00.txt")
	os.Stdout = writer

	Exec()
	writer.Close()
	defer reader.Close()
	actualResult, _ := ioutil.ReadAll(reader)
	actualResult = bytes.TrimSpace(actualResult)

	assert.Equal(t, expectedResult, actualResult)
}

func TestSumAllPrimeNumberOfInput10ShouldReturn17(t *testing.T) {
	expectedSumPrimeNumber := 17

	actualSumPrimeNumber := SumAllPrimeNumberOf(10)

	assert.Equal(t, expectedSumPrimeNumber, actualSumPrimeNumber)
}

func TestSumAllPrimeNumberOfInput5ShouldReturn10(t *testing.T) {
	expectedSumPrimeNumber := 10

	actualSumPrimeNumber := SumAllPrimeNumberOf(5)

	assert.Equal(t, expectedSumPrimeNumber, actualSumPrimeNumber)
}

func TestIsPrimeInput4ShouldReturnFalse(t *testing.T) {
	expectedIsPrime := false

	actualIsPrime := IsPrime(4)

	assert.Equal(t, expectedIsPrime, actualIsPrime)
}

func TestIsPrimeInput3ShouldReturnTrue(t *testing.T) {
	expectedIsPrime := true

	actualIsPrime := IsPrime(3)

	assert.Equal(t, expectedIsPrime, actualIsPrime)
}

func TestIsPrimeInput2ShouldReturnTrue(t *testing.T) {
	expectedIsPrime := true

	actualIsPrime := IsPrime(2)

	assert.Equal(t, expectedIsPrime, actualIsPrime)
}

func TestIsPrimeInput1ShouldReturnFalse(t *testing.T) {
	expectedIsPrime := false

	actualIsPrime := IsPrime(1)

	assert.Equal(t, expectedIsPrime, actualIsPrime)
}
