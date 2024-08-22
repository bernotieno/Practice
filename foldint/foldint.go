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
	printDigits(result)
	z01.PrintRune('\n')
}

func printDigits(n int) {
	if n < 10 {
		z01.PrintRune('0' + rune(n))
		return
	}

	printDigits(n / 10)
	z01.PrintRune('0' + rune(n%10))
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
