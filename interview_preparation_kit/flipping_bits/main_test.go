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

func TestMainInputFile10ShouldBeOutputFile10(t *testing.T) {
	expectedResult, _ := ioutil.ReadFile("output/output10.txt")
	expectedResult = bytes.TrimSpace(expectedResult)
	cmd := exec.Command("go", "run", "main.go")
	cmd.Stdin, _ = os.Open("input/input10.txt")

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

func TestGetFlippingBitsInput4294967295ShouldReturn0(t *testing.T) {
	expectedFlippingBits := uint32(0)

	actualFlippingBits := GetFlippingBits(4294967295)

	assert.Equal(t, expectedFlippingBits, actualFlippingBits)
}

func TestGetFlippingBitsInput0ShouldReturn4294967295(t *testing.T) {
	expectedFlippingBits := uint32(4294967295)

	actualFlippingBits := GetFlippingBits(0)

	assert.Equal(t, expectedFlippingBits, actualFlippingBits)
}
