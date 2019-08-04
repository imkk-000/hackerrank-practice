package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var primeCache []bool

func SieveOfEratosthenes(n int) {
	primeCache = make([]bool, n+1)
	primeCache[1] = true
	for i := 2; i <= n; i++ {
		primeCache[i] = false
	}
	for p := 2; p*p <= n; p++ {
		if !primeCache[p] {
			for i := p * p; i <= n; i += p {
				primeCache[i] = true
			}
		}
	}
}

func IsPrime(number int) bool {
	return !primeCache[number]
}

func SumAllPrimeNumberOf(number int) int {
	sum := 2
	for i := 3; i <= number; i += 2 {
		if IsPrime(i) {
			sum += i
		}
	}
	return sum
}

func Exec() {
	SieveOfEratosthenes(1299743)
	reader := bufio.NewReader(os.Stdin)
	t, _ := strconv.Atoi(readLine(reader))
	for i := 0; i < t; i++ {
		n, _ := strconv.Atoi(readLine(reader))
		fmt.Println(SumAllPrimeNumberOf(n))
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
