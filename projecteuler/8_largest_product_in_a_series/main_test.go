package main_test

import (
	"bytes"
	. "dojo"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLargestProductOfInput2709360626ShouldReturn0(t *testing.T) {
	expectedLargestProduct := 0

	actualLargestProduct := FindLargestProductOf("2709360626", 10, 5)

	assert.Equal(t, expectedLargestProduct, actualLargestProduct)
}

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
	actualResult = bytes.TrimSpace(expectedResult)

	assert.Equal(t, expectedResult, actualResult)
}

func TestFindLargestProductOfInput3675356291ShouldReturn3150(t *testing.T) {
	expectedLargestProduct := 3150

	actualLargestProduct := FindLargestProductOf("3675356291", 10, 5)

	assert.Equal(t, expectedLargestProduct, actualLargestProduct)
}

func TestFindProductOfInput67535ShouldReturn3150(t *testing.T) {
	expectedProduct := 3150

	actualProduct := FindProductOf([]int{6, 7, 5, 3, 5})

	assert.Equal(t, expectedProduct, actualProduct)
}

func TestConvertToIntsInput2709360626ShouldReturnInts2709360626(t *testing.T) {
	expectedInts := []int{2, 7, 0, 9, 3, 6, 0, 6, 2, 6}

	actualInts := ConvertToInts("2709360626")

	assert.Equal(t, expectedInts, actualInts)
}

func TestConvertToIntsInput3675356291ShouldReturnInts3675356291(t *testing.T) {
	expectedInts := []int{3, 6, 7, 5, 3, 5, 6, 2, 9, 1}

	actualInts := ConvertToInts("3675356291")

	assert.Equal(t, expectedInts, actualInts)
}
