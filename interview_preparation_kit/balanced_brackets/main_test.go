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

func TestMainInputFile00ShouldBeOutputFile00(t *testing.T) {
	expectedResult, _ := ioutil.ReadFile("output/output00.txt")
	expectedResult = bytes.TrimSpace(expectedResult)
	cmd := exec.Command("go", "run", "main.go")
	cmd.Stdin, _ = os.Open("input/input00.txt")

	actualResult, _ := cmd.Output()
	actualResult = bytes.TrimSpace(actualResult)

	assert.Equal(t, expectedResult, actualResult)
}

func TestMainInputFile20ShouldBeOutputFile20(t *testing.T) {
	expectedResult, _ := ioutil.ReadFile("output/output20.txt")
	expectedResult = bytes.TrimSpace(expectedResult)
	cmd := exec.Command("go", "run", "main.go")
	cmd.Stdin, _ = os.Open("input/input20.txt")

	actualResult, _ := cmd.Output()
	actualResult = bytes.TrimSpace(actualResult)

	assert.Equal(t, expectedResult, actualResult)
}

func TestMainInputFile19ShouldBeOutputFile19(t *testing.T) {
	expectedResult, _ := ioutil.ReadFile("output/output19.txt")
	expectedResult = bytes.TrimSpace(expectedResult)
	cmd := exec.Command("go", "run", "main.go")
	cmd.Stdin, _ = os.Open("input/input19.txt")

	actualResult, _ := cmd.Output()
	actualResult = bytes.TrimSpace(actualResult)

	assert.Equal(t, expectedResult, actualResult)
}

func TestMainInputFile18ShouldBeOutputFile18(t *testing.T) {
	expectedResult, _ := ioutil.ReadFile("output/output18.txt")
	expectedResult = bytes.TrimSpace(expectedResult)
	cmd := exec.Command("go", "run", "main.go")
	cmd.Stdin, _ = os.Open("input/input18.txt")

	actualResult, _ := cmd.Output()
	actualResult = bytes.TrimSpace(actualResult)

	assert.Equal(t, expectedResult, actualResult)
}

func TestIsBalancedBracketsInputBrackets2ShouldReturnTrue(t *testing.T) {
	actualIsBalancedBrackets := IsBalancedBrackets("{[()]}")

	assert.True(t, actualIsBalancedBrackets)
}

func TestIsBalancedBracketsInputBrackets1ShouldReturnFalse(t *testing.T) {
	actualIsBalancedBrackets := IsBalancedBrackets("{[(])}")

	assert.False(t, actualIsBalancedBrackets)
}
