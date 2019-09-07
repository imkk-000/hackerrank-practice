package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func AlternatingCharacters(chars string) (counter int) {
	lengthOfChars := len(chars)
	prevChar := chars[0]
	for i := 1; i < lengthOfChars; i++ {
		if prevChar == chars[i] {
			counter++
		} else {
			prevChar = chars[i]
		}
	}
	return
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	t, _ := strconv.Atoi(readLine(reader))
	for i := 0; i < t; i++ {
		fmt.Println(AlternatingCharacters(readLine(reader)))
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimSpace(string(str))
}
