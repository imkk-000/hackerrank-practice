package main_test

import (
	"bytes"
	. "dojo"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecInputFile04ShouldBeOutputFile04(t *testing.T) {
	expectedResult, _ := ioutil.ReadFile("output/output04.txt")
	os.Stdin, _ = os.Open("input/input04.txt")
	reader, writer, _ := os.Pipe()
	os.Stdout = writer

	Exec()
	writer.Close()
	defer reader.Close()
	actualResult, _ := ioutil.ReadAll(reader)
	actualResult = bytes.TrimSpace(actualResult)

	assert.Equal(t, expectedResult, actualResult)
}

func TestFindEvenFibonacciInput3ShouldBe34(t *testing.T) {
	expectedEvenFibonacci := int64(34)

	actualEvenFibonacci := FindEvenFibonacci(3)

	assert.Equal(t, expectedEvenFibonacci, actualEvenFibonacci)
}

func TestFindEvenFibonacciInput2ShouldBe8(t *testing.T) {
	expectedEvenFibonacci := int64(8)

	actualEvenFibonacci := FindEvenFibonacci(2)

	assert.Equal(t, expectedEvenFibonacci, actualEvenFibonacci)
}

func TestFindEvenFibonacciInput1ShouldBe2(t *testing.T) {
	expectedEvenFibonacci := int64(2)

	actualEvenFibonacci := FindEvenFibonacci(1)

	assert.Equal(t, expectedEvenFibonacci, actualEvenFibonacci)
}
