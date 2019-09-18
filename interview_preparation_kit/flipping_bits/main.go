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

func GetFlippingBits(number uint32) uint32 {
	return math.MaxUint32 - number
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	t, _ := strconv.Atoi(readLine(reader))
	for i := 0; i < t; i++ {
		number, _ := strconv.ParseUint(readLine(reader), 10, 32)
		fmt.Println(GetFlippingBits(uint32(number)))
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimSpace(string(str))
}
