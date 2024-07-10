package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		n := Atoi(os.Args[1])
		if n == 0 {
			printer("00000000")
		}

		res := printBits(n, "01")
		printer(res)
	}
}

func printBits(num int, base string) string {
	result := ""

	for num > 0 {
		rem := num % len(base)
		result = string(base[rem]) + result
		num /= len(base)
	}
	return format(result)
}

func format(result string) string {
	zeros := ""

	if len(result) >= 8 {
		return result
	}

	for i := 0; i < 8-len(result); i++ {
		zeros += "0"
	}

	return zeros + result
}

func Atoi(str string) int {
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

func printer(str string) {
	for _, char := range str {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
