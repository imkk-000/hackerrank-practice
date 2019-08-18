package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func StepToSeaLevelFlag(step rune) int {
	if step == 'D' {
		return -1
	}
	return 1
}

func uphill(step rune) bool {
	return step == 'U'
}

func downhill(step rune) bool {
	return step == 'D'
}

func CountValley(sequenceStep string) (valleyCounter int) {
	var seaLevel int
	for _, step := range sequenceStep {
		seaLevel += StepToSeaLevelFlag(step)
		if seaLevel == 0 && uphill(step) {
			valleyCounter++
		}
	}
	return
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	readLine(reader)
	fmt.Println(CountValley(readLine(reader)))
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimSpace(string(str))
}
