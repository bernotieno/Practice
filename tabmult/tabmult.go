package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}
	isNeg := false
	if n < 0 {
		isNeg = true
		n = -n
	}
	var number []rune
	for n > 0 {
		digit := n % 10
		number = append([]rune{rune(digit + '0')}, number...)
		n /= 10
	}
	if isNeg {
		number = append([]rune{'-'}, number...)
	}
	return string(number)
}

func Atoi(str string) int {
	num := 0
	sign := 1

	for i, char := range str {
		if i == 0 && char == '-' {
			sign = -1
			continue
		} else if i == 0 && char == '+' {
			continue
		}
		if char < '0' || char > '9' {
			return 0
		}
		num = num*10 + int(char-'0')
	}
	return num * sign
}

func Tabmult(s string) {
	for i := 1; i <= 9; i++ {
		STR := Itoa(i) + " " + "X" + " " + s + " " + "=" + " " + Itoa(i*Atoi(s))

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
