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

func TestMainInputFile07ShouldBeOutputFile07(t *testing.T) {
	expectedResult, _ := ioutil.ReadFile("output/output07.txt")
	expectedResult = bytes.TrimSpace(expectedResult)
	cmd := exec.Command("go", "run", "main.go")
	cmd.Stdin, _ = os.Open("input/input07.txt")

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

func TestMainInputFile00ShouldBeOutputFile00(t *testing.T) {
	expectedResult, _ := ioutil.ReadFile("output/output00.txt")
	expectedResult = bytes.TrimSpace(expectedResult)
	cmd := exec.Command("go", "run", "main.go")
	cmd.Stdin, _ = os.Open("input/input00.txt")

	actualResult, _ := cmd.Output()
	actualResult = bytes.TrimSpace(actualResult)

	assert.Equal(t, expectedResult, actualResult)
}

func TestIsCommonSubStringInputhiAndworldShouldReturnFalse(t *testing.T) {
	actualIsCommonSubstring := IsCommonSubString("hi", "world")

	assert.False(t, actualIsCommonSubstring)
}

func TestIsCommonSubStringInputhelloAndworldShouldReturnTrue(t *testing.T) {
	actualIsCommonSubstring := IsCommonSubString("hello", "world")

	assert.True(t, actualIsCommonSubstring)
}
