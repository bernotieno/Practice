package functions

import (
	"os"

	"github.com/01-edu/z01"
)

func isDigit(r rune) bool {
	return '0' <= r && r <= '9' || r == '-' || r == '+'
}

func ParseInt(s string) (int64, bool) {
	var result int64
	var neg bool
	for _, char := range s {
		if char == '-' {
			neg = !neg
		} else if isDigit(char) {
			result = result*10 + int64(char-'0')
		} else {
			return 0, false // Invalid input
		}
	}
	if neg {
		result = -result
	}
	return result, true
}

func Doop() {
	if len(os.Args) != 4 {
		return
	}

	val1, err1 := ParseInt(os.Args[1])
	operator := os.Args[2]
	val2, err2 := ParseInt(os.Args[3])

	if !err1 || !err2 {
		return
	}

	var result int64

	switch operator {
	case "+":
		result = val1 + val2
	case "-":
		result = val1 - val2
	case "*":
		result = val1 * val2
	case "/":
		if val2 == 0 {
			z01.PrintRune('N')
			z01.PrintRune('o')
			z01.PrintRune(' ')
			z01.PrintRune('d')
			z01.PrintRune('i')
			z01.PrintRune('v')
			z01.PrintRune('i')
			z01.PrintRune('s')
			z01.PrintRune('i')
			z01.PrintRune('o')
			z01.PrintRune('n')
			z01.PrintRune(' ')
			z01.PrintRune('b')
			z01.PrintRune('y')
			z01.PrintRune(' ')
			z01.PrintRune('0')
			z01.PrintRune('\n')
			return
		}
		result = val1 / val2
	case "%":
		if val2 == 0 {
			z01.PrintRune('N')
			z01.PrintRune('o')
			z01.PrintRune(' ')
			z01.PrintRune('m')
			z01.PrintRune('o')
			z01.PrintRune('d')
			z01.PrintRune('u')
			z01.PrintRune('l')
			z01.PrintRune('o')
			z01.PrintRune(' ')
			z01.PrintRune('b')
			z01.PrintRune('y')
			z01.PrintRune(' ')
			z01.PrintRune('0')
			z01.PrintRune('\n')
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

func itoa(n int64) string {
	var result string
	var neg bool

	if n < 0 {
		neg = true
		n = -n
	}

	for n > 0 {
		digit := n % 10
		result = string(digit+'0') + result
		n /= 10
	}

	if neg {
		result = "-" + result
	}

	return result
}
