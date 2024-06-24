package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	num, err := strconv.Atoi(os.Args[1])
	if err != nil || num <= 0 {
		return
	}

	factors := primeFactors(num)
	if len(factors) == 0 {
		return
	}

	for i, factor := range factors {
		for _, n := range Itoa(factor) {
			z01.PrintRune(n)
		}
		if i < len(factors)-1 {
			z01.PrintRune('*')
		}
	}
	z01.PrintRune('\n')
}

// Function to find prime factors of a number
func primeFactors(n int) []int {
	factors := make([]int, 0)
	for i := 2; i <= n; i++ {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
	}
	return factors
}

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	var digits []rune
	for n > 0 {
		digit := n % 10
		digits = append([]rune{rune('0' + digit)}, digits...)
		n /= 10
	}
	return string(digits)
}
