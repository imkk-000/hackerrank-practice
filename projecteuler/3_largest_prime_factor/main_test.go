package main_test

import (
	"bytes"
	. "dojo"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecInputFile01ShouldBeOutputFile01(t *testing.T) {
	expectedResult, _ := ioutil.ReadFile("output/output01.txt")
	expectedResult = bytes.TrimSpace(expectedResult)
	os.Stdin, _ = os.Open("input/input01.txt")
	reader, writer, _ := os.Pipe()
	os.Stdout = writer

	Exec()
	writer.Close()
	defer reader.Close()
	actualResult, _ := ioutil.ReadAll(reader)
	actualResult = bytes.TrimSpace(actualResult)

	assert.Equal(t, expectedResult, actualResult)
}

func TestExecInputFile00ShouldBeOutputFile00(t *testing.T) {
	expectedResult, _ := ioutil.ReadFile("output/output00.txt")
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

func TestGetLargestPrimeOfFactorInput1FactorsShouldBe0(t *testing.T) {
	expectedLargestPrime := 0

	actualLargestPrime := GetLargestPrimeOfFactor(1)

	assert.Equal(t, expectedLargestPrime, actualLargestPrime)
}

func TestGetLargestPrimeOfFactorInput13195ShouldBe29(t *testing.T) {
	expectedLargestPrime := 29

	actualLargestPrime := GetLargestPrimeOfFactor(13195)

	assert.Equal(t, expectedLargestPrime, actualLargestPrime)
}

func TestGetLargestPrimeOfFactorInput9ShouldBeX(t *testing.T) {
	expectedLargestPrime := 3

	actualLargestPrime := GetLargestPrimeOfFactor(9)

	assert.Equal(t, expectedLargestPrime, actualLargestPrime)
}

func TestGetLargestPrimeOfFactorInput17ShouldBe17(t *testing.T) {
	expectedLargestPrime := 17

	actualLargestPrime := GetLargestPrimeOfFactor(17)

	assert.Equal(t, expectedLargestPrime, actualLargestPrime)
}

func TestGetLargestPrimeOfFactorInput10ShouldBe5(t *testing.T) {
	expectedLargestPrime := 5

	actualLargestPrime := GetLargestPrimeOfFactor(10)

	assert.Equal(t, expectedLargestPrime, actualLargestPrime)
}

func TestGetFactorOfInput13195ShouldBeFactors(t *testing.T) {
	expectedFactors := []int{5, 7, 13, 29}

	actualFactors := GetFactorOf(13195)

	assert.Equal(t, expectedFactors, actualFactors)
}

func TestGetFactorOfInput17ShouldBeFactors(t *testing.T) {
	expectedFactors := []int{17}

	actualFactors := GetFactorOf(17)

	assert.Equal(t, expectedFactors, actualFactors)
}

func TestGetFactorOfInput10ShouldBeFactors(t *testing.T) {
	expectedFactors := []int{2, 5}

	actualFactors := GetFactorOf(10)

	assert.Equal(t, expectedFactors, actualFactors)
}

func TestIsPrimeInput4ShouldBeFalse(t *testing.T) {
	expectedIsPrime := false

	actualIsPrime := IsPrime(4)

	assert.Equal(t, expectedIsPrime, actualIsPrime)
}

func TestIsPrimeInput3ShouldBeTrue(t *testing.T) {
	expectedIsPrime := true

	actualIsPrime := IsPrime(3)

	assert.Equal(t, expectedIsPrime, actualIsPrime)
}

func TestIsPrimeInput2ShouldBeTrue(t *testing.T) {
	expectedIsPrime := true

	actualIsPrime := IsPrime(2)

	assert.Equal(t, expectedIsPrime, actualIsPrime)
}

func TestIsPrimeInput1ShouldBeFalse(t *testing.T) {
	expectedIsPrime := false

	actualIsPrime := IsPrime(1)

	assert.Equal(t, expectedIsPrime, actualIsPrime)
}
