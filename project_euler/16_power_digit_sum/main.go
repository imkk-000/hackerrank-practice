package main

import (
	"fmt"
	"math/big"
)

func GetPowerOf(n int64) string {
	return new(big.Int).Exp(big.NewInt(2), big.NewInt(n), nil).String()
}

func SumOfPowerDigit(n int64) (sum int) {
	for _, digit := range GetPowerOf(n) {
		sum += int(digit - '0')
	}
	return sum
}

func main() {
	var t, n int64
	fmt.Scanf("%d", &t)
	for i := int64(0); i < t; i++ {
		fmt.Scanf("%d", &n)
		fmt.Println(SumOfPowerDigit(n))
	}
}
