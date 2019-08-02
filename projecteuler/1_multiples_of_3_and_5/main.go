package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func SumOf(number int) int64 {
	var sum int64
	var i int
	for i = 3; i < number; i += 3 {
		if i%15 == 0 {
			continue
		}
		sum += int64(i)
	}
	for i = 5; i < number; i += 5 {
		sum += int64(i)
	}
	return sum
}

func Exec() {
	reader := bufio.NewReader(os.Stdin)
	t, _ := strconv.Atoi(ReadLine(reader))
	for i := 0; i < t; i++ {
		n, _ := strconv.Atoi(ReadLine(reader))
		fmt.Println(SumOf(n))
	}
}

func ReadLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(str), "\r\n")
}

func main() {
	Exec()
}
