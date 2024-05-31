package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	value1 := Atoi(os.Args[1])
	operator := os.Args[2]
	value2 := Atoi(os.Args[3])

	if (value1 >= 9223372036854775807 || value1 <= -9223372036854775807) || (value2 >= 9223372036854775807 || value2 <= -9223372036854775807) {
		return
	}

	var result int
	switch operator {
	case "+":
		result = value1 + value2
	case "-":
		result = value1 - value2
	case "*":
		result = value1 * value2
	case "/":
		if value2 == 0 {
			word := "No division by 0"
			Print(word)
			return
		}
		result = value1 / value2
	case "%":
		if value2 == 0 {
			word := "No modulo by 0"
			Print(word)
			return
		}
		result = value1 % value2
	default:
		return
	}
	r := Itoa(int(result))
	Print(r)
}

func Print(str string) {
	for _, char := range str {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}

func Itoa(n int) string {
	if n == 0 {
		return "0"
	} else if n < 0 {
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

func Atoi(s string) int {
	var number int
	sign := 1

	for idx, char := range s {
		if char == '-' && idx == 0 {
			sign = -1
		} else if char == '+' && idx == 0 {
			sign = 1
		} else if char >= '0' && char <= '9' {
			number = number*10 + int(char-'0')
		} else {
			return 0
		}
	}
	return number * sign
}
