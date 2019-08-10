package main

import (
	"bufio"
	"fmt"
	"io"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func Add(firstNumber, secondNumber string) string {
	firstNumberBig, secondNumberBig, resultBig := new(big.Int), new(big.Int), new(big.Int)
	firstNumberBig, _ = firstNumberBig.SetString(firstNumber, 10)
	secondNumberBig, _ = secondNumberBig.SetString(secondNumber, 10)
	resultBig = resultBig.Add(firstNumberBig, secondNumberBig)
	return resultBig.String()
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	t, _ := strconv.Atoi(readLine(reader))
	sum := "0"
	for i := 0; i < t; i++ {
		sum = Add(sum, readLine(reader))
	}
	fmt.Println(sum[:10])
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimSpace(string(str))
}
