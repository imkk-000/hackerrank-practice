package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func Sum1ToN(number int64) int64 {
	return (number*number + number) / 2
}

func SumOf(number int64) int64 {
	number--
	return (Sum1ToN(number/3) * 3) +
		(Sum1ToN(number/5) * 5) -
		(Sum1ToN(number/15) * 15)
}

func Exec() {
	reader := bufio.NewReader(os.Stdin)
	t, _ := strconv.Atoi(ReadLine(reader))
	for i := 0; i < t; i++ {
		n, _ := strconv.Atoi(ReadLine(reader))
		fmt.Println(SumOf(int64(n)))
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
