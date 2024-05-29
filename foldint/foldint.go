package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func FoldInt(f func(int, int) int, a []int, n int) {
	var tol int
	for _, value := range a {
		tol += value
	}
	result := f(tol, n)
	for _, char := range Itoa(result) {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
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
	var res []byte
	for nb > 0 {
		digit := nb % 10
		res = append([]byte{byte('0' + digit)}, res...)
		nb /= 10
	}
	if isNeg {
		res = append([]byte{'-'}, res...)
	}
	return string(res)
}

func Add(a, b int) int {
	return a + b
}

func Mul(a, b int) int {
	return a * b
}

func Sub(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func main() {
	table := []int{1, 2, 3}
	ac := 93
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
	fmt.Println()

	table = []int{0}
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
}
