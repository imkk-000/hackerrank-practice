package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const markChar = '-'

func MakeAnagram(firstWord, secondWord string) int {
	countMap := make(map[rune]int)
	for _, ch := range firstWord {
		countMap[ch]++
	}
	for _, ch := range secondWord {
		countMap[ch]--
	}
	var countRemove int
	for ch := 'a'; ch <= 'z'; ch++ {
		if count, existing := countMap[ch]; existing {
			if count < 0 {
				count *= -1
			}
			countRemove += count
		}
	}
	return countRemove
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	firstWord := readLine(reader)
	secondWord := readLine(reader)
	fmt.Println(MakeAnagram(firstWord, secondWord))
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimSpace(string(str))
}
