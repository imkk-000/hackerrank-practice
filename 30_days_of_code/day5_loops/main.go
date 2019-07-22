package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func GetMultiplicationTable(number int) []int {
	multiplicationTable := make([]int, 10)
	for i := 1; i <= 10; i++ {
		multiplicationTable[i-1] = number * i
	}
	return multiplicationTable
}

func printMultiplicationTable(number int, result []int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", number, i, result[i-1])
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int(nTemp)
	printMultiplicationTable(n, GetMultiplicationTable(n))
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
