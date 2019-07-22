package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func PlusNumber(firstNumber, secondNumber uint64) string {
	return fmt.Sprint(firstNumber + secondNumber)
}

func PlusFloatNumber(firstFloatNumber, secondFloatNumber float64) string {
	return fmt.Sprintf("%.1f", firstFloatNumber+secondFloatNumber)
}

func ConcatenateString(firstString, secondString string) string {
	return firstString + secondString
}

func main() {
	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)

	// Declare second integer, double, and String variables.
	var sumNumber uint64
	var sumFloatNumber float64
	var text string

	// Read and save an integer, double, and String to your variables.
	scanner.Scan()
	sumNumber, _ = strconv.ParseUint(scanner.Text(), 10, 64)
	scanner.Scan()
	sumFloatNumber, _ = strconv.ParseFloat(scanner.Text(), 64)
	scanner.Scan()
	text = scanner.Text()

	// Print the sum of both integer variables on a new line.
	fmt.Println(PlusNumber(i, sumNumber))

	// Print the sum of the double variables on a new line.
	fmt.Println(PlusFloatNumber(d, sumFloatNumber))

	// Concatenate and print the String variables on a new line
	// The 's' variable above should be printed first.
	fmt.Println(ConcatenateString(s, text))
}
