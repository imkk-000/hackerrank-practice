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

func TestMainInputFile11ShouldBeOutputFile11(t *testing.T) {
	expectedResult, _ := ioutil.ReadFile("output/output11.txt")
	expectedResult = bytes.TrimSpace(expectedResult)
	cmd := exec.Command("go", "run", "main.go")
	cmd.Stdin, _ = os.Open("input/input11.txt")

	actualResult, _ := cmd.Output()
	actualResult = bytes.TrimSpace(actualResult)

	assert.Equal(t, expectedResult, actualResult)
}

func TestMainInputFile06ShouldBeOutputFile06(t *testing.T) {
	expectedResult, _ := ioutil.ReadFile("output/output06.txt")
	expectedResult = bytes.TrimSpace(expectedResult)
	cmd := exec.Command("go", "run", "main.go")
	cmd.Stdin, _ = os.Open("input/input06.txt")

	actualResult, _ := cmd.Output()
	actualResult = bytes.TrimSpace(actualResult)

	assert.Equal(t, expectedResult, actualResult)
}

func TestMainInputFile04ShouldBeOutputFile04(t *testing.T) {
	expectedResult, _ := ioutil.ReadFile("output/output04.txt")
	expectedResult = bytes.TrimSpace(expectedResult)
	cmd := exec.Command("go", "run", "main.go")
	cmd.Stdin, _ = os.Open("input/input04.txt")

	actualResult, _ := cmd.Output()
	actualResult = bytes.TrimSpace(actualResult)

	assert.Equal(t, expectedResult, actualResult)
}

func TestFibonacciNumberInput30ShouldReturn832040(t *testing.T) {
	expectedNumber := 832040

	actualNumber := FibonacciNumber(30)

	assert.Equal(t, expectedNumber, actualNumber)
}

func TestFibonacciNumberInput3ShouldReturn2(t *testing.T) {
	expectedNumber := 2

	actualNumber := FibonacciNumber(3)

	assert.Equal(t, expectedNumber, actualNumber)
}

func TestFibonacciNumberInput2ShouldReturn1(t *testing.T) {
	expectedNumber := 1

	actualNumber := FibonacciNumber(2)

	assert.Equal(t, expectedNumber, actualNumber)
}

func TestFibonacciNumberInput1ShouldReturn1(t *testing.T) {
	expectedNumber := 1

	actualNumber := FibonacciNumber(1)

	assert.Equal(t, expectedNumber, actualNumber)
}

func TestFibonacciNumberInput0ShouldReturn0(t *testing.T) {
	expectedNumber := 0

	actualNumber := FibonacciNumber(0)

	assert.Equal(t, expectedNumber, actualNumber)
}
