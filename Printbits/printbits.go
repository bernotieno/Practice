package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Printbits(nb int) {
	for i := 7; i >= 0; i-- {
		bin := (nb >> i) & 1
		z01.PrintRune(rune(bin) + '0')
	}
	z01.PrintRune('\n')
}

func Atoi(str string) int {
	// var result int

	// for _, digit := range str {
	// 	if digit < '0' || digit > '9' {
	// 		return 0
	// 	}
	// 	result = result*10 + int(digit-'0')
	// }
	// return result

	result := 0
	sign := 1

	for i, ch := range str {
		if ch < '0' || ch > '9' {
			return 0
		}
		if i == 0 && ch == '-' {
			sign = -1
			continue
		} else if i == 0 && ch == '+' {
			continue
		}
		result *= 10
		result += int(ch - '0')
	}
	return result * sign
}

func main() {
	if len(os.Args) != 2 {
		return
	}
	args := Atoi(os.Args[1])
	Printbits(args)
}
