package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func CountRepeatedString(word string, n int) (count int) {
	lengthWord := len(word)
	var countFirstCharcter int
	for i := 0; i < lengthWord; i++ {
		if word[i] == 'a' {
			countFirstCharcter++
		}
	}
	count += countFirstCharcter * (n / lengthWord)
	for i := 0; i < n%lengthWord; i++ {
		if word[i] == 'a' {
			count++
		}
	}
	return
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	word := readLine(reader)
	n, _ := strconv.Atoi(readLine(reader))
	fmt.Println(CountRepeatedString(word, n))
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimSpace(string(str))
}
