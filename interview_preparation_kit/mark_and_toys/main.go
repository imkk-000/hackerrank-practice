package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func GetMaximumToys(k int, toys []int) (count int) {
	sort.Ints(toys)
	// cut numbers are more than K
	var sum int
	for _, toy := range toys {
		if sum+toy > k {
			break
		}
		sum += toy
		count++
	}
	return count
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	var k int
	tempProp := strings.Split(readLine(reader), " ")
	k, _ = strconv.Atoi(tempProp[1])
	tempToys := strings.Split(readLine(reader), " ")
	var toys []int
	for _, tempToy := range tempToys {
		toy, _ := strconv.Atoi(tempToy)
		toys = append(toys, toy)
	}
	fmt.Println(GetMaximumToys(k, toys))
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimSpace(string(str))
}
