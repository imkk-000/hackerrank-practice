package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func ConvertToInts(str string) (Ints []int) {
	var n int
	for _, s := range strings.Split(str, "") {
		n, _ = strconv.Atoi(s)
		Ints = append(Ints, n)
	}
	return
}

func FindProductOf(numbers []int) int {
	product := 1
	for _, n := range numbers {
		product *= n
	}
	return product
}

func FindLargestProductOf(numberString string, N, K int) int {
	var max, product int
	numbers := ConvertToInts(numberString)
	for i := 0; i <= N-K; i++ {
		product = FindProductOf(numbers[i : K+i])
		if max < product {
			max = product
		}
	}
	return max
}

func Exec() {
	reader := bufio.NewReader(os.Stdin)
	t, _ := strconv.Atoi(readLine(reader))
	for i := 0; i < t; i++ {
		nTemp := strings.Split(readLine(reader), " ")
		N, _ := strconv.Atoi(nTemp[0])
		K, _ := strconv.Atoi(nTemp[1])
		largestProduct := FindLargestProductOf(readLine(reader), N, K)
		fmt.Println(largestProduct)
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
