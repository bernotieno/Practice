package main

import (
	"os"

	"github.com/01-edu/z01"
)

const (
	maxInt = 9223372036854775807
	minInt = -9223372036854775808
)

func Atoi(s string) int {
	result := 0
	sign := 1

	for i, char := range s {
		if i == 0 && char == '-' {
			sign = -1
			continue
		} else if i == 0 && char == '+' {
			continue
		}
		if char < '0' || char > '9' {
			return 0
		}

		digit := int(char - '0')

		// Check for overflow before multiplying and adding
		if result > (maxInt-digit)/10 {
			if sign == 1 {
				return maxInt
			} else {
				return minInt
			}
		}
		result = result*10 + digit
	}

	return result * sign
}

func Doop() {
	if len(os.Args) != 4 {
		return
	}

	val1 := Atoi(os.Args[1])
	operator := os.Args[2]
	val2 := Atoi(os.Args[3])

	var result int

	switch operator {
	case "+":
		if (val2 > 0 && val1 > maxInt-val2) || (val2 < 0 && val1 < minInt-val2) {
			return
		}
		result = val1 + val2
	case "-":
		if (val2 < 0 && val1 > maxInt+val2) || (val2 > 0 && val1 < minInt+val2) {
			return
		}
		result = val1 - val2
	case "*":
		if val2 != 0 && ((val1 > maxInt/val2 && val2 > 0) || (val1 < minInt/val2 && val2 < 0)) {
			return
		}

		result = val1 * val2
	case "/":
		if val2 == 0 {
			return
		}
		result = val1 / val2
	case "%":
		if val2 == 0 {
			return
		}
		result = val1 % val2
	default:
		return
	}

	for _, char := range itoa(result) {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}

	var result string
	var neg bool

	if n < 0 {
		neg = true
		n = -n
	}

	for n > 0 {
		digit := n % 10
		result = string(rune(digit+'0')) + result
		n /= 10
	}

	if neg {
		result = "-" + result
	}

	return result
}

func main() {
	Doop()
}
