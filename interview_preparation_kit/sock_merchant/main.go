package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func CountPairNumber(sockMerchantList []int) (counter int) {
	var lookUpTable = make([]bool, 101)
	for _, sockMerchant := range sockMerchantList {
		if lookUpTable[sockMerchant] {
			counter++
		}
		lookUpTable[sockMerchant] = !lookUpTable[sockMerchant]
	}
	return
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	strconv.Atoi(readLine(reader))
	var sockMerchantList []int
	var n int
	for _, sockMerchant := range strings.Split(readLine(reader), " ") {
		n, _ = strconv.Atoi(sockMerchant)
		sockMerchantList = append(sockMerchantList, n)
	}
	fmt.Println(CountPairNumber(sockMerchantList))
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimSpace(string(str))
}
