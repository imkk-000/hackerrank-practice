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

func TestMainInputFile10ShouldBeOutputFile10(t *testing.T) {
	expectedResult, _ := ioutil.ReadFile("output/output10.txt")
	expectedResult = bytes.TrimSpace(expectedResult)
	cmd := exec.Command("go", "run", "main.go")
	cmd.Stdin, _ = os.Open("input/input10.txt")

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

func TestShiftLoopInput12345D4ShouldReturn51234(t *testing.T) {
	expectedSlice := []int{5, 1, 2, 3, 4}

	actualSlice := ShiftLoop(4, []int{1, 2, 3, 4, 5})

	assert.Equal(t, expectedSlice, actualSlice)
}

func TestShiftLeftRotateInput12345ShouldReturn23451(t *testing.T) {
	expectedSlice := []int{2, 3, 4, 5, 1}

	actualSlice := ShiftLeftRotate([]int{1, 2, 3, 4, 5})

	assert.Equal(t, expectedSlice, actualSlice)
}
