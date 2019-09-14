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

func TestMainInputFile17ShouldBeOutputFile17(t *testing.T) {
	expectedResult, _ := ioutil.ReadFile("output/output17.txt")
	expectedResult = bytes.TrimSpace(expectedResult)
	cmd := exec.Command("go", "run", "main.go")
	cmd.Stdin, _ = os.Open("input/input17.txt")

	actualResult, _ := cmd.Output()
	actualResult = bytes.TrimSpace(actualResult)

	assert.Equal(t, expectedResult, actualResult)
}

func TestMainInputFile16ShouldBeOutputFile16(t *testing.T) {
	expectedResult, _ := ioutil.ReadFile("output/output16.txt")
	expectedResult = bytes.TrimSpace(expectedResult)
	cmd := exec.Command("go", "run", "main.go")
	cmd.Stdin, _ = os.Open("input/input16.txt")

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

func TestGetMaximumToysInputK5000AndToysListsShoudReturn7(t *testing.T) {
	expectedMaximumToys := 7

	actualMaximumToys := GetMaximumToys(5000, []int{1, 12, 5, 111, 200, 1000, 10})

	assert.Equal(t, expectedMaximumToys, actualMaximumToys)
}

func TestGetMaximumToysInputK1AndToysListsShoudReturn0(t *testing.T) {
	expectedMaximumToys := 0

	actualMaximumToys := GetMaximumToys(1, []int{12, 5, 111, 200, 1000, 10})

	assert.Equal(t, expectedMaximumToys, actualMaximumToys)
}

func TestGetMaximumToysInputK50AndToysListsShoudReturn4(t *testing.T) {
	expectedMaximumToys := 4

	actualMaximumToys := GetMaximumToys(50, []int{1, 12, 5, 111, 200, 1000, 10})

	assert.Equal(t, expectedMaximumToys, actualMaximumToys)
}
