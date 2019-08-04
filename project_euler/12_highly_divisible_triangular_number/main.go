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

func FindFactorOf(number int) (factors []int) {
	for i := 1; i <= int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			factors = append(factors, i)
			if i == number/i {
				continue
			}
			factors = append(factors, number/i)
		}
	}
	return
}

func SumOf(number int) int {
	return (number*number + number) / 2
}

type IntsSlice [][]int

func (p IntsSlice) Len() int           { return len(p) }
func (p IntsSlice) Less(i, j int) bool { return p[i][0] < p[j][0] }
func (p IntsSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func FindHiglyDivisibleTriangularNumber(n int) int {
	var factorCache [][]int
	var sumN = SumOf(n)
	if n == 1 {
		sumN = 2
	}
	for i := 2; i <= sumN; i++ {
		sumByIndex := SumOf(i)
		factorCache = append(factorCache, []int{len(FindFactorOf(sumByIndex)) - 1, sumByIndex})
	}
	sort.Sort(IntsSlice(factorCache))
	for i := 0; i < len(factorCache); i++ {
		if factorCache[i][0] >= n {
			return factorCache[i][1]
		}
	}
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	t, _ := strconv.Atoi(readLine(reader))
	for i := 0; i < t; i++ {
		n, _ := strconv.Atoi(readLine(reader))
		fmt.Println(FindHiglyDivisibleTriangularNumber(n))
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimSpace(string(str))
}
