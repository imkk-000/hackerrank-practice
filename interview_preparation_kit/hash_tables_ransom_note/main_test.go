package main_test

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"

	. "dojo"

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

func TestMainInputFile20ShouldBeOutputFile20(t *testing.T) {
	expectedResult, _ := ioutil.ReadFile("output/output20.txt")
	expectedResult = bytes.TrimSpace(expectedResult)
	cmd := exec.Command("go", "run", "main.go")
	cmd.Stdin, _ = os.Open("input/input20.txt")

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

func TestCheckMagazineInputStringTwoListsShouldReturnFalse(t *testing.T) {
	magazine := []string{"two", "times", "three", "is", "not", "four"}
	note := []string{"two", "times", "two", "is", "four"}

	actualCheckMagazine := CheckMagazine(magazine, note)

	assert.False(t, actualCheckMagazine)
}

func TestCheckMagazineInputStringTwoListsShouldReturnTrue(t *testing.T) {
	magazine := []string{"give", "me", "one", "grand", "today", "night"}
	note := []string{"give", "one", "grand", "today"}

	actualCheckMagazine := CheckMagazine(magazine, note)

	assert.True(t, actualCheckMagazine)
}
