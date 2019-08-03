package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func SumSquareOf(n int64) int64 {
	return ((n*n - 1) * (3*n + 2) * n) / 12
}

func Exec() {
	reader := bufio.NewReader(os.Stdin)
	t, _ := strconv.Atoi(readLine(reader))
	for i := 0; i < t; i++ {
		n, _ := strconv.ParseInt(readLine(reader), 10, 64)
		fmt.Println(SumSquareOf(n))
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimSpace(string(str))
}

func main() {
	Exec()
}
