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

func TestFindCollatzSequenceStepWithMultipleTestCase(t *testing.T) {
	testCases := []struct {
		input, expect int
	}{
		{input: 1, expect: 1},
		{input: 2, expect: 2},
		{input: 3, expect: 8},
		{input: 4, expect: 3},
	}

	for _, testCase := range testCases {
		actual := FindCollatzSequenceStep(testCase.input)

		assert.Equal(t, testCase.expect, actual)
	}
}
