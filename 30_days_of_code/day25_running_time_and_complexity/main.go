package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func FindPrimeNumber(number int) bool {
	if number == 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int(nTemp)
	for i := 0; i < n; i++ {
		nTemp, err = strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		if FindPrimeNumber(int(nTemp)) {
			fmt.Println("Prime")
		} else {
			fmt.Println("Not prime")
		}
	}
}

func readLine(reader *bufio.Reader) string {
	s, _, err := reader.ReadLine()
	if err != nil {
		return ""
	}
	return strings.TrimRight(string(s), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
