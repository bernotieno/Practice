package main

import "fmt"

func Itoa(nb int) string {
	if nb == 0 {
		return "0"
	}
	isNeg := false

	if nb < 0 {
		isNeg = true
		nb = -nb
	}
	var result []rune
	for nb > 0 {
		digit := nb % 10
		result = append([]rune{rune('0' + digit)}, result...)
		nb /= 10
	}
	if isNeg {
		result = append([]rune{'-'}, result...)
	}
	return string(result)
}

func main() {
	fmt.Println(Itoa(12345))
	fmt.Println(Itoa(0))
	fmt.Println(Itoa(-1234))
	fmt.Println(Itoa(987654321))
}
