package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func IsCommonSubString(firstString, secondString string) bool {
	lengthFirstString := len(firstString)
	lengthSecondString := len(secondString)
	if lengthFirstString < lengthSecondString {
		firstString, secondString = secondString, firstString
		lengthFirstString, lengthSecondString = lengthSecondString, lengthFirstString
	}
	commonSubString := make(map[rune]bool)
	for _, ch := range firstString {
		commonSubString[ch] = true
	}
	for _, ch := range secondString {
		if _, common := commonSubString[ch]; common {
			return true
		}
	}
	return false
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	t, _ := strconv.Atoi(readLine(reader))
	for i := 0; i < t; i++ {
		temp := readLine(reader)
		if IsCommonSubString(temp, readLine(reader)) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimSpace(string(str))
}
