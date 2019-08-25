package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func CheckMagazine(magazine, note []string) bool {
	hashTables := make(map[string]int)
	for _, str := range magazine {
		hashTables[str]++
	}
	for _, str := range note {
		if _, existing := hashTables[str]; !existing {
			return false
		}
		hashTables[str]--
		if hashTables[str] <= 0 {
			delete(hashTables, str)
		}
	}
	return true
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	readLine(reader)
	magazine := strings.Split(readLine(reader), " ")
	note := strings.Split(readLine(reader), " ")
	if CheckMagazine(magazine, note) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimSpace(string(str))
}
