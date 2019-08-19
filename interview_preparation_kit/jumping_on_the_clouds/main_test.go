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

func TestFindMinJumpInput000010ShouldReturn4(t *testing.T) {
	expectedMinJump := 3

	actualMinJump := FindMinJump([]int{0, 0, 0, 0, 1, 0})

	assert.Equal(t, expectedMinJump, actualMinJump)
}

func TestFindMinJumpInput0010010ShouldReturn4(t *testing.T) {
	expectedMinJump := 4

	actualMinJump := FindMinJump([]int{0, 0, 1, 0, 0, 1, 0})

	assert.Equal(t, expectedMinJump, actualMinJump)
}
