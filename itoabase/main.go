package main

import (
	"fmt"
)

func ItoaBase(value, base int) string {
	if base < 2 || base > 16 {
		return ""
	}

	if value == 0 {
		return "0"
	}

	digits := "0123456789ABCDEF"
	result := ""
	isNegative := value < 0
	if isNegative {
		value = -value
	}

	for value > 0 {
		remainder := value % base
		result = string(digits[remainder]) + result
		value /= base
	}

	if isNegative {
		result = "-" + result
	}

	return result
}

func main() {
	fmt.Println(ItoaBase(255, 16))
	fmt.Println(ItoaBase(-255, 16))
	fmt.Println(ItoaBase(10, 2))
	fmt.Println(ItoaBase(-10, 2))
	fmt.Println(ItoaBase(0, 10))
}
