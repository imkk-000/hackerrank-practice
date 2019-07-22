package main

import (
	"fmt"
	"strings"
)

const (
	ODD  = 1
	EVEN = 0
)

func GetStringFromIndexed(indexType int, text string) string {
	var newText strings.Builder
	for i, r := range text {
		if i%2 == indexType {
			newText.WriteRune(r)
		}
	}
	return newText.String()
}

func main() {
	var t int
	var text string
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		fmt.Scanln(&text)
		fmt.Println(GetStringFromIndexed(EVEN, text) + " " + GetStringFromIndexed(ODD, text))
	}
}
