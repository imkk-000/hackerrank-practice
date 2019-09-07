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

func TestMainInputFile14ShouldBeOutputFile14(t *testing.T) {
	expectedResult, _ := ioutil.ReadFile("output/output14.txt")
	expectedResult = bytes.TrimSpace(expectedResult)
	cmd := exec.Command("go", "run", "main.go")
	cmd.Stdin, _ = os.Open("input/input14.txt")

	actualResult, _ := cmd.Output()
	actualResult = bytes.TrimSpace(actualResult)

	assert.Equal(t, expectedResult, actualResult)
}

func TestMainInputFile13ShouldBeOutputFile13(t *testing.T) {
	expectedResult, _ := ioutil.ReadFile("output/output13.txt")
	expectedResult = bytes.TrimSpace(expectedResult)
	cmd := exec.Command("go", "run", "main.go")
	cmd.Stdin, _ = os.Open("input/input13.txt")

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

func TestAlternatingCharactersInputABBAShouldReturn1(t *testing.T) {
	expectedAlterChar := 1

	actualAlterChar := AlternatingCharacters("ABBA")

	assert.Equal(t, expectedAlterChar, actualAlterChar)
}

func TestAlternatingCharactersInputABABShouldReturn0(t *testing.T) {
	expectedAlterChar := 0

	actualAlterChar := AlternatingCharacters("ABAB")

	assert.Equal(t, expectedAlterChar, actualAlterChar)
}

func TestAlternatingCharactersInputAABBShouldReturn2(t *testing.T) {
	expectedAlterChar := 2

	actualAlterChar := AlternatingCharacters("AABB")

	assert.Equal(t, expectedAlterChar, actualAlterChar)
}

func TestAlternatingCharactersInputAShouldReturn0(t *testing.T) {
	expectedAlterChar := 0

	actualAlterChar := AlternatingCharacters("A")

	assert.Equal(t, expectedAlterChar, actualAlterChar)
}
