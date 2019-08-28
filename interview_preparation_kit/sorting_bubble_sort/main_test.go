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

func TestMainInputFile03ShouldBeOutputFile03(t *testing.T) {
	expectedResult, _ := ioutil.ReadFile("output/output03.txt")
	expectedResult = bytes.TrimSpace(expectedResult)
	cmd := exec.Command("go", "run", "main.go")
	cmd.Stdin, _ = os.Open("input/input03.txt")

	actualResult, _ := cmd.Output()
	actualResult = bytes.TrimSpace(actualResult)

	assert.Equal(t, expectedResult, actualResult)
}

func TestMainInputFile01ShouldBeOutputFile01(t *testing.T) {
	expectedResult, _ := ioutil.ReadFile("output/output01.txt")
	expectedResult = bytes.TrimSpace(expectedResult)
	cmd := exec.Command("go", "run", "main.go")
	cmd.Stdin, _ = os.Open("input/input01.txt")

	actualResult, _ := cmd.Output()
	actualResult = bytes.TrimSpace(actualResult)

	assert.Equal(t, expectedResult, actualResult)
}

func TestMainInputFile00ShouldBeOutputFile00(t *testing.T) {
	expectedResult, _ := ioutil.ReadFile("output/output00.txt")
	expectedResult = bytes.TrimSpace(expectedResult)
	cmd := exec.Command("go", "run", "main.go")
	cmd.Stdin, _ = os.Open("input/input00.txt")

	actualResult, _ := cmd.Output()
	actualResult = bytes.TrimSpace(actualResult)

	assert.Equal(t, expectedResult, actualResult)
}

func TestBubbleSortInputList1234ShouldReturnList123AndSwapCount0(t *testing.T) {
	expectedLists := []int{1, 2, 3, 4}
	expectedSwapCount := 0

	actualLists, actualSwapCount := BubbleSort([]int{1, 2, 3, 4})

	assert.Equal(t, expectedLists, actualLists)
	assert.Equal(t, expectedSwapCount, actualSwapCount)
}

func TestBubbleSortInputList321ShouldReturnList123AndSwapCount3(t *testing.T) {
	expectedLists := []int{1, 2, 3}
	expectedSwapCount := 3

	actualLists, actualSwapCount := BubbleSort([]int{3, 2, 1})

	assert.Equal(t, expectedLists, actualLists)
	assert.Equal(t, expectedSwapCount, actualSwapCount)
}
