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

func TestSumSquareOfInput10ShouldReturn2640(t *testing.T) {
	expectedNumber := int64(2640)

	actualNumber := SumSquareOf(10)

	assert.Equal(t, expectedNumber, actualNumber)
}

func TestSumSquareOfInput4ShouldReturn70(t *testing.T) {
	expectedNumber := int64(70)

	actualNumber := SumSquareOf(4)

	assert.Equal(t, expectedNumber, actualNumber)
}
