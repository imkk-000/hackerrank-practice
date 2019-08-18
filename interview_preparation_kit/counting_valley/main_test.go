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

func TestMainInputFile02ShouldReturnOutputFile02(t *testing.T) {
	expectedResult, _ := ioutil.ReadFile("output/output02.txt")
	expectedResult = bytes.TrimSpace(expectedResult)
	cmd := exec.Command("go", "run", "main.go")
	cmd.Stdin, _ = os.Open("input/input02.txt")

	actualResult, _ := cmd.Output()
	actualResult = bytes.TrimSpace(actualResult)

	assert.Equal(t, expectedResult, actualResult)
}

func TestMainInputFile01ShouldReturnOutputFile01(t *testing.T) {
	expectedResult, _ := ioutil.ReadFile("output/output01.txt")
	expectedResult = bytes.TrimSpace(expectedResult)
	cmd := exec.Command("go", "run", "main.go")
	cmd.Stdin, _ = os.Open("input/input01.txt")

	actualResult, _ := cmd.Output()
	actualResult = bytes.TrimSpace(actualResult)

	assert.Equal(t, expectedResult, actualResult)
}

func TestMainInputFile00ShouldReturnOutputFile00(t *testing.T) {
	expectedResult, _ := ioutil.ReadFile("output/output00.txt")
	expectedResult = bytes.TrimSpace(expectedResult)
	cmd := exec.Command("go", "run", "main.go")
	cmd.Stdin, _ = os.Open("input/input00.txt")

	actualResult, _ := cmd.Output()
	actualResult = bytes.TrimSpace(actualResult)

	assert.Equal(t, expectedResult, actualResult)
}

func TestCountValleyInputDDUUUUDDShouldReturn2(t *testing.T) {
	expectedCountValley := 1

	actualCountValley := CountValley("DDUUUUDD")

	assert.Equal(t, expectedCountValley, actualCountValley)
}

func TestCountValleyInputUDDDUDUUShouldReturn1(t *testing.T) {
	expectedCountValley := 1

	actualCountValley := CountValley("UDDDUDUU")

	assert.Equal(t, expectedCountValley, actualCountValley)
}

func TestStepToSeaLevelFlagInputUShouldReturnMinus1(t *testing.T) {
	expectedSeaLevelFlag := -1

	actualSeaLevelFlag := StepToSeaLevelFlag('D')

	assert.Equal(t, expectedSeaLevelFlag, actualSeaLevelFlag)
}

func TestStepToSeaLevelFlagInputUShouldReturn1(t *testing.T) {
	expectedSeaLevelFlag := 1

	actualSeaLevelFlag := StepToSeaLevelFlag('U')

	assert.Equal(t, expectedSeaLevelFlag, actualSeaLevelFlag)
}
