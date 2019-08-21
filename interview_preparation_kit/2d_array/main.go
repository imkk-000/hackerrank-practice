package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// distance
const d = 3

func FindMaximumSum(matrix [][]int) int {
	lengthMatrix := len(matrix)
	max := -100
	var sum int
	for m := 0; m < lengthMatrix-d+1; m++ {
		for n := 0; n < lengthMatrix-d+1; n++ {
			for i := m; i < m+d; i++ {
				for j := n; j < n+d; j++ {
					sum += matrix[i][j]
				}
			}
			sum -= (matrix[m+1][n] + matrix[m+1][n+2])
			if max < sum {
				max = sum
			}
			sum = 0
		}
	}
	return max
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	var arr [][]int
	for i := 0; i < 6; i++ {
		arrRowTemp := strings.Split(readLine(reader), " ")
		var arrRow []int
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, _ := strconv.ParseInt(arrRowItem, 10, 64)
			arrItem := int(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}
		arr = append(arr, arrRow)
	}
	fmt.Println(FindMaximumSum(arr))
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimSpace(string(str))
}
