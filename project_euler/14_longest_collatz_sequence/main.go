package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func FindCollatzSequenceStep(N int) int {
	var step = 1
	for N != 1 {
		if N < cacheLimit {
			if caching[N] > 0 {
				return caching[N] + step
			}
		}
		if N%2 == 0 {
			N = N / 2
		} else {
			N = 3*N + 1
		}
		step++
	}
	return step
}

func GetLongestCollatzSequenceStep(N int) int {
	var maxStep int
	var maxN = 1
	for i := N; i >= 1; i-- {
		if maxStep < caching[i] {
			maxStep = caching[i]
			maxN = i
		}
	}
	return maxN
}

const cacheLimit = 5000001

var caching = make([]int, cacheLimit)

func main() {
	for i := 2; i < cacheLimit; i++ {
		caching[i] = FindCollatzSequenceStep(i)
	}

	reader := bufio.NewReader(os.Stdin)
	t, _ := strconv.Atoi(readLine(reader))
	for i := 0; i < t; i++ {
		n, _ := strconv.Atoi(readLine(reader))
		fmt.Println(GetLongestCollatzSequenceStep(n))
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimSpace(string(str))
}
