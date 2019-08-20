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

func TestMainInputFile21ShouldBeOutputFile21(t *testing.T) {
	expectedResult, _ := ioutil.ReadFile("output/output21.txt")
	expectedResult = bytes.TrimSpace(expectedResult)
	cmd := exec.Command("go", "run", "main.go")
	cmd.Stdin, _ = os.Open("input/input21.txt")

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

func TestCountRepeatedStringInputaN1000000000000ShouldReturn(t *testing.T) {
	expectedCountRepeatedString := 6

	actualCountRepeatedString := CountRepeatedString("a", 6)

	assert.Equal(t, expectedCountRepeatedString, actualCountRepeatedString)
}

func TestCountRepeatedStringInputaN1000000000000ShouldReturn1000000000000(t *testing.T) {
	expectedCountRepeatedString := 1000000000000

	actualCountRepeatedString := CountRepeatedString("a", 1000000000000)

	assert.Equal(t, expectedCountRepeatedString, actualCountRepeatedString)
}

func TestCountRepeatedStringInputabaN10ShouldReturn7(t *testing.T) {
	expectedCountRepeatedString := 7

	actualCountRepeatedString := CountRepeatedString("aba", 10)

	assert.Equal(t, expectedCountRepeatedString, actualCountRepeatedString)
}
