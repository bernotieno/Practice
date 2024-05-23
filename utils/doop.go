package functions

import (
	"os"

	"github.com/01-edu/z01"
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
		result = result*10 + int(char-'0')
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

	if val1 < -9223372036854775808 || val1 > 9223372036854775807 {
		return
	}

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

func itoa(n int) string {
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

// package functions

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// func Doop() {
// 	if len(os.Args) != 4 {
// 		return
// 	}

// 	val1, err1 := strconv.ParseInt(os.Args[1], 10, 64)
// 	operator := os.Args[2]
// 	val2, err2 := strconv.ParseInt(os.Args[3], 10, 64)

// 	if err1 != nil || err2 != nil {
// 		return
// 	}

// 	var result int64

// 	switch operator {
// 	case "+":
// 		result = val1 + val2
// 	case "-":
// 		result = val1 - val2
// 	case "*":
// 		result = val1 * val2
// 	case "/":
// 		if val2 == 0 {
// 			fmt.Println("No division by 0")
// 			return
// 		}
// 		result = val1 / val2
// 	case "%":
// 		if val2 == 0 {
// 			fmt.Println("No modulo by 0")
// 			return
// 		}
// 		result = val1 % val2
// 	default:
// 		return
// 	}

// 	fmt.Println(result)
// }
