package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func BubbleSort(list []int) ([]int, int) {
	var swapCount int
	lengthList := len(list)
	for i := 0; i < lengthList; i++ {
		for j := 0; j < lengthList-1; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
				swapCount++
			}
		}
	}
	return list, swapCount
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	n, _ := strconv.Atoi(readLine(reader))
	temp := strings.Split(readLine(reader), " ")
	list := make([]int, n)
	for i, tempNumber := range temp {
		list[i], _ = strconv.Atoi(tempNumber)
	}
	sorting, swapCount := BubbleSort(list)
	fmt.Println("Array is sorted in", swapCount, "swaps.")
	fmt.Println("First Element:", sorting[0])
	fmt.Println("Last Element:", sorting[len(sorting)-1])
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimSpace(string(str))
}
