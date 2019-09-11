package main

import "fmt"

func FibonacciNumber(number int) int {
	switch number {
	case 0:
		return 0
	case 1:
		return 1
	}
	return FibonacciNumber(number-1) + FibonacciNumber(number-2)
}

func main() {
	var number int
	fmt.Scanf("%d", &number)
	fmt.Println(FibonacciNumber(number))
}
