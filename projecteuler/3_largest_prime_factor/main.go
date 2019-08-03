package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
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

func GetFactorOf(number int) (factors []int) {
	for i := 1; i <= int(math.Sqrt(float64(number))); i++ {
		if number%i != 0 {
			continue
		}
		if IsPrime(i) {
			factors = append(factors, i)
		}
		if i == number/i {
			continue
		}
		if IsPrime(number / i) {
			factors = append(factors, number/i)
		}
	}
	sort.Ints(factors)
	return
}

func GetLargestPrimeOfFactor(number int) int {
	factors := GetFactorOf(number)
	factorsLength := len(factors)
	if factorsLength > 0 {
		return factors[factorsLength-1]
	}
	return 0
}

func Exec() {
	reader := bufio.NewReader(os.Stdin)
	t, _ := strconv.Atoi(readLine(reader))
	for i := 0; i < t; i++ {
		n, _ := strconv.Atoi(readLine(reader))
		n = GetLargestPrimeOfFactor(n)
		fmt.Println(n)
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
