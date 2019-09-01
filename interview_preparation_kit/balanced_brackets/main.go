package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var bracketMap = map[rune]rune{
	'}': '{',
	']': '[',
	')': '(',
}

func IsBalancedBrackets(brackets string) bool {
	var stack []rune
	for _, bracket := range brackets {
		if bracketPair, existing := bracketMap[bracket]; existing {
			if len(stack) > 0 {
				n := len(stack) - 1
				if bracketPair != stack[n] {
					return false
				}
				stack = stack[:n]
			} else {
				return false
			}
		} else {
			stack = append(stack, bracket)
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	t, _ := strconv.Atoi(readLine(reader))
	for i := 0; i < t; i++ {
		if IsBalancedBrackets(readLine(reader)) {
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
