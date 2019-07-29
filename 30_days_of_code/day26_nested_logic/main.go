package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Date struct {
	Day, Month, Year int
}

func GetDate(dateString string) (date Date) {
	fmt.Sscanf(dateString, "%d %d %d", &date.Day, &date.Month, &date.Year)
	return
}

func CalculateFine(actualDate, expectedDate Date) int {
	switch {
	case expectedDate.Year > actualDate.Year:
		return 0
	case expectedDate.Year < actualDate.Year:
		return 10000
	}
	switch {
	case expectedDate.Month > actualDate.Month:
		return 0
	case expectedDate.Month < actualDate.Month:
		return 500 * (actualDate.Month - expectedDate.Month)
	}
	switch {
	case expectedDate.Day > actualDate.Day:
		return 0
	case expectedDate.Day < actualDate.Day:
		return 15 * (actualDate.Day - expectedDate.Day)
	}
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	actualDate := GetDate(readLine(reader))
	expectedDate := GetDate(readLine(reader))
	fmt.Println(CalculateFine(actualDate, expectedDate))
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
