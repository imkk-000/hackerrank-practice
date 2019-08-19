package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func FindMinJump(list []int) (jumpCounter int) {
	lengthList := len(list)
	for i := 0; i < lengthList-1; i++ {
		switch {
		case i+2 < lengthList && list[i+2] != 1:
			i++
			jumpCounter++
		case list[i+1] != 1:
			jumpCounter++
		}
	}
	return
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	readLine(reader)
	tempList := strings.Split(readLine(reader), " ")
	var list []int
	for _, elm := range tempList {
		n, _ := strconv.Atoi(elm)
		list = append(list, n)
	}
	fmt.Println(FindMinJump(list))
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimSpace(string(str))
}
