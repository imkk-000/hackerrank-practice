package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func CalculateTotalCost(mealCost float64, tipPercent, taxPercent int32) string {
	tip := mealCost * float64(tipPercent) / 100
	tax := mealCost * float64(taxPercent) / 100
	return fmt.Sprintf("%.0f", mealCost+tip+tax)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	meal_cost, err := strconv.ParseFloat(readLine(reader), 64)
	checkError(err)

	tip_percentTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	tip_percent := int32(tip_percentTemp)

	tax_percentTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	tax_percent := int32(tax_percentTemp)

	fmt.Println(CalculateTotalCost(meal_cost, tip_percent, tax_percent))
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
