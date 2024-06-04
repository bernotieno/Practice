package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Atoi(str string) int {
	result := 0
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
		result = result*10 + int(char-'0')
	}
	return result * sign
}

func Itoa(nb int) string {
	if nb == 0 {
		return "0"
	}
	isNeg := false
	if nb < 0 {
		isNeg = true
		nb = -nb
	}
	var number []rune
	for nb > 0 {
		digit := nb % 10
		number = append([]rune{rune(digit + '0')}, number...)
		nb /= 10
	}
	if isNeg {
		number = append([]rune{'-'}, number...)
	}

	return string(number)
}

func main() {
	num1 := Atoi(os.Args[1])
	num2 := Atoi(os.Args[2])

	gcd := 0

	if num1 > num2 {
		for i := 1; i <= num1; i++ {
			if num1%i == 0 && num2%i == 0 {
				gcd = i
			}
		}
	}
	if num2 > num1 {
		for i := 1; i <= num2; i++ {
			if num1%i == 0 && num2%i == 0 {
				gcd = i
			}
		}
	}

	Number := Itoa(gcd)
	Print(Number)
}

func Print(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
