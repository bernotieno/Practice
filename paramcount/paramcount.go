package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Itoa(n int) string {
	// If n is zero
	if n == 0 {
		return "0"
	}

	// Negative number
	isNeg := false
	if n < 0 {
		isNeg = true
		n = -n
	}

	// Convert to string
	var result []byte
	for n > 0 {
		digit := n % 10
		result = append([]byte{byte('0' + digit)}, result...)
		n /= 10
	}
	if isNeg {
		result = append([]byte{'-'}, result...)
	}
	return string(result)
}

func main() {
	var count int

	arguments := os.Args[1:]

	for range arguments {
		count++
	}

	for _, val := range Itoa(count) {
		z01.PrintRune(val)
	}
	z01.PrintRune('\n')
}
