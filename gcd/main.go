package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	num1 := Atoi(os.Args[1])
	num2 := Atoi(os.Args[2])

	gcd := 0

	if num1 > num2 {
		for i := 1; i <= num1; i++ {
			if num1%i == 0 && num2%i == 0 {
				gcd = i
			}
		}
	} else if num2 > num1 {
		for i := 1; i <= num2; i++ {
			if num1%i == 0 && num2%i == 0 {
				gcd = i
			}
		}
	}
	number := Itoa(gcd)
	for _, c := range number {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
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

/*func Gcd(a, b int) int {
	if a != 0 {
		a, b = b, a%b
	}
	return a
}*/
