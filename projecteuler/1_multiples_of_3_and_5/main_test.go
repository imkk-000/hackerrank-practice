package main_test

import (
	"bytes"
	. "dojo"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecInputFile05ShouldBeOutputFile05(t *testing.T) {
	expectedResult, _ := ioutil.ReadFile("output/output05.txt")
	os.Stdin, _ = os.Open("input/input05.txt")
	reader, writer, _ := os.Pipe()
	os.Stdout = writer

	Exec()
	writer.Close()
	defer reader.Close()
	actualResult, _ := ioutil.ReadAll(reader)
	actualResult = bytes.TrimSpace(actualResult)

	assert.Equal(t, expectedResult, actualResult)
}

func TestSumOfInput100ShouldBeNumber2318(t *testing.T) {
	expectedNumber := int64(2318)

	actualNumber := SumOf(100)

	assert.Equal(t, expectedNumber, actualNumber)
}

func TestSumOfInput10ShouldBeNumber23(t *testing.T) {
	expectedNumber := int64(23)

	actualNumber := SumOf(10)

	assert.Equal(t, expectedNumber, actualNumber)
}
