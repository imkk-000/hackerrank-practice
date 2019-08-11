package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const cacheLimit = 5000000

var cachingStep = make([]int, cacheLimit+1)
var cachingLongest = make([]int, cacheLimit+1)

func LoadCache(limit int) {
	var maxStep, maxN int
	for i := 1; i <= limit; i++ {
		cachingStep[i] = FindCollatzSequenceStep(i)
		if maxStep <= cachingStep[i] {
			maxStep = cachingStep[i]
			maxN = i
		}
		cachingLongest[i] = maxN
	}
}

func FindCollatzSequenceStep(N int) (step int) {
	step = 1
	for N != 1 {
		if N < cacheLimit {
			if cachingStep[N] > 0 {
				return cachingStep[N] + step - 1
			}
		}
		if N%2 == 0 {
			N = N / 2
		} else {
			N = 3*N + 1
		}
		step++
	}
	return
}

func GetLongestCollatzSequenceStep(N int) int {
	return cachingLongest[N]
}

func main() {
	LoadCache(cacheLimit)

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
