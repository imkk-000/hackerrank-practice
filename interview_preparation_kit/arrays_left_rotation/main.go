package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func ShiftLeftRotate(arr []int) []int {
	arr = append(arr, arr[0])
	return arr[1:]
}

func ShiftLoop(n int, arr []int) []int {
	for i := 0; i < n; i++ {
		arr = ShiftLeftRotate(arr)
	}
	return arr
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	tempArr := strings.Split(readLine(reader), " ")
	n, _ := strconv.Atoi(tempArr[0])
	d, _ := strconv.Atoi(tempArr[1])
	tempArr = strings.Split(readLine(reader), " ")
	arr := []int{}
	for i := 0; i < n; i++ {
		elm, _ := strconv.Atoi(tempArr[i])
		arr = append(arr, elm)
	}
	for _, elm := range ShiftLoop(d, arr) {
		fmt.Print(elm, " ")
	}
	fmt.Println()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimSpace(string(str))
}
