package functions

import (
	"fmt"
)

// FoldInt applies the given function f to each element of the slice and the accumulator,
// prints the intermediate result after each application, and returns the final accumulated value.
func FoldInt(f func(int, int) int, slice []int, acc int) {
	for _, value := range slice {
		acc = f(acc, value)
		fmt.Println(acc)
	}
}

// Add function for addition.
func Add(a, b int) int {
	return a + b
}

// Sub function for subtraction.
func Sub(a, b int) int {
	return a - b
}

// Mul function for multiplication.
func Mul(a, b int) int {
	return a * b
}
