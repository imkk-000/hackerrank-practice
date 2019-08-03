package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

var primeCache = map[int]bool{
	1: false,
}

func IsPrime(number int) bool {
	if isPrime, existing := primeCache[number]; existing {
		return isPrime
	}
	for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			primeCache[number] = false
			break
		}
	}
	if _, existing := primeCache[number]; !existing {
		primeCache[number] = true
	}
	return primeCache[number]
}

func FindPrimeNumberFromIndexOf(number int) (prime int) {
	if number == 1 {
		prime = 2
	}
	for i, count := 3, 1; count < number; i += 2 {
		if IsPrime(i) {
			prime = i
			count++
		}
	}
	return
}

func Exec() {
	reader := bufio.NewReader(os.Stdin)
	t, _ := strconv.Atoi(readLine(reader))
	for i := 0; i < t; i++ {
		n, _ := strconv.Atoi(readLine(reader))
		fmt.Println(FindPrimeNumberFromIndexOf(n))
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
