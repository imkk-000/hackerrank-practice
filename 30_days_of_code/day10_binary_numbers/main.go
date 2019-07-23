package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func ConvertNumberFromDecimalToBinary(decimalNumber int) string {
	return strconv.FormatInt(int64(decimalNumber), 2)
}

func FindMaxCountBitFromBinaryNumber(binaryNumber string) int {
	var max, count int
	for _, b := range binaryNumber {
		if b == '1' {
			count++
		} else {
			count = 0
		}
		if max < count {
			max = count
		}
	}
	return max
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int(nTemp)
	binaryNumber := ConvertNumberFromDecimalToBinary(n)
	fmt.Println(FindMaxCountBitFromBinaryNumber(binaryNumber))
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
