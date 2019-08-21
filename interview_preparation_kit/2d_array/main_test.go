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

func TestMainInputFile08ShouldBeOutputFile08(t *testing.T) {
	expectedResult, _ := ioutil.ReadFile("output/output08.txt")
	expectedResult = bytes.TrimSpace(expectedResult)
	cmd := exec.Command("go", "run", "main.go")
	cmd.Stdin, _ = os.Open("input/input08.txt")

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

func TestFindMaximumSumInputMatrixShouldReturn19(t *testing.T) {
	expectedMaximumSum := 19
	inputData := [][]int{
		[]int{1, 1, 1, 0, 0, 0},
		[]int{0, 1, 0, 0, 0, 0},
		[]int{1, 1, 1, 0, 0, 0},
		[]int{0, 0, 2, 4, 4, 0},
		[]int{0, 0, 0, 2, 0, 0},
		[]int{0, 0, 1, 2, 4, 0},
	}

	actaulMaximumSum := FindMaximumSum(inputData)

	assert.Equal(t, expectedMaximumSum, actaulMaximumSum)
}
