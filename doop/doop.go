package main

import (
	"os"
)

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}
	sign := ""
	if n < 0 {
		sign = "-"
		n = -n
	}
	q := ""
	for n > 0 {
		digits := n % 10
		q = string(rune('0'+digits)) + q
		n /= 10
	}
	return sign + q
}

func Atoi(s string) (int, bool) {
	sign := 1
	q := 0

	for i, v := range s {
		if v == '-' && i == 0 {
			sign = -1
		} else if v == '+' && i == 0 {
			sign = 1
		} else if v >= '0' && v <= '9' {
			q = q*10 + int(v-'0')
		} else {
			return 0, false
		}
	}
	return sign * q, true
}

func PrintStr(s string) {
	os.Stdout.WriteString(s)
	os.Stdout.WriteString("\n")
}

func Calculation(a, operator, b string) string {
	value1, ok1 := Atoi(a)
	value2, ok2 := Atoi(b)

	if !ok1 || !ok2 {
		return ""
	}

	// Define max int64 value to handle large numbers
	const maxInt64 = 9223372036854775807

	if (value1 >= maxInt64 || value1 <= -maxInt64) || (value2 >= maxInt64 || value2 <= -maxInt64) {
		return ""
	}

	q := 0
	switch operator {
	case "+":
		q = value1 + value2
	case "-":
		q = value1 - value2
	case "*":
		q = value1 * value2
	case "/":
		if value2 == 0 {
			return "No division by 0"
		}
		q = value1 / value2
	case "%":
		if value2 == 0 {
			return "No modulo by 0"
		}
		q = value1 % value2
	default:
		return ""
	}
	return Itoa(q)
}

func main() {
	if len(os.Args) != 4 {
		return
	}
	value1, operator, value2 := os.Args[1], os.Args[2], os.Args[3]

	result := Calculation(value1, operator, value2)

	if result != "" {
		PrintStr(result)
	}
}
