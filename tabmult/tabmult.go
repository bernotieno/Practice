package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Itoa(nb int) string {
	if nb == 0 {
		return "0"
	}
	isNeg := false
	if nb < 0 {
		isNeg = true
		nb = -nb
	}
	var digits []rune
	for nb > 0 {
		digit := nb % 10
		digits = append([]rune{rune('0' + digit)}, digits...)
		nb /= 10
	}
	if isNeg {
		digits = append([]rune{'-'}, digits...)
	}
	return string(digits)
}

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
		result = result*10 + int(char-'0')
	}
	return result * sign
}

func Tabmult(str string) {
	for i := 1; i <= 9; i++ {
		STR := Itoa(i) + " " + "X" + " " + str + " " + "=" + " " + Itoa(i*Atoi(str))

		for _, char := range STR {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}

func main() {
	if len(os.Args) != 2 {
		return
	}
	args := os.Args[1]
	Tabmult(args)
}
