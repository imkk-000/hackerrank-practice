package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var fiboCache = map[int64]int64{
	1: 2,
	2: 8,
}

func FindEvenFibonacci(number int64) int64 {
	if fibo, existing := fiboCache[number]; existing {
		return fibo
	}
	fiboCache[number] = 4*FindEvenFibonacci(number-1) + FindEvenFibonacci(number-2)
	return fiboCache[number]
}

func Exec() {
	reader := bufio.NewReader(os.Stdin)
	t, _ := strconv.Atoi(readLine(reader))
	for i := 0; i < t; i++ {
		nTemp, _ := strconv.Atoi(readLine(reader))
		var sum int64
		var r int64
		for j := 1; ; j++ {
			r = FindEvenFibonacci(int64(j))
			if r > int64(nTemp) {
				break
			}
			sum += r
		}
		fmt.Println(sum)
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(str), "\r\n")
}

func main() {
	Exec()
}
