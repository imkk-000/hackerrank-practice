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

func TestAdd37107287533902102798797998220837590246510135740250And46376937677490009712648124896970078050417018260538ShouldBe83484225211392112511446123117807668296927154000788(t *testing.T) {
	expectedBigInt := "83484225211392112511446123117807668296927154000788"

	actualBigInt := Add("37107287533902102798797998220837590246510135740250", "46376937677490009712648124896970078050417018260538")

	assert.Equal(t, expectedBigInt, actualBigInt)
}
