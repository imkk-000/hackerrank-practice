package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func BubbleSortWithSwapTime(arr []int) ([]int, int) {
	var swapTime int
	arrLength := len(arr)
	for i := 0; i < arrLength; i++ {
		for j := 0; j < arrLength-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapTime++
			}
		}
	}
	return arr, swapTime
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 32)
	checkError(err)

	var arr []int
	for _, arrStringTemp := range strings.Split(readLine(reader), " ") {
		nTemp, err = strconv.ParseInt(arrStringTemp, 10, 32)
		checkError(err)
		arr = append(arr, int(nTemp))
	}

	arr, swapTimes := BubbleSortWithSwapTime(arr)
	fmt.Printf("Array is sorted in %d swaps.\n", swapTimes)
	fmt.Printf("First Element: %d\n", arr[0])
	fmt.Printf("Last Element: %d\n", arr[len(arr)-1])
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
