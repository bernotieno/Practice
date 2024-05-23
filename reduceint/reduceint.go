// Overall, this program showcases how to implement a generic reduction function in Go,
// allowing for different reduction operations to be applied to a slice of integers,
// with the flexibility to print the result in a custom format using the z01 package.

package main

import "github.com/01-edu/z01"

func printDigits(n int) {
	// base case just one digit
	if n < 10 {
		z01.PrintRune('0' + rune(n))
		return
	}

	// case 2 more than 1 digit
	// recursion to print the digits except the last one
	printDigits(n / 10)
	z01.PrintRune('0' + rune(n%10))
}

func ReduceInt(a []int, f func(int, int) int) {
	// Initialize the accumulator with the first element of the slice
	acc := a[0]

	// Iterate over the rest of the elements in the slice
	for _, val := range a[1:] {
		// Apply the provided function to the accumulator and the current element
		acc = f(acc, val)
	}

	// Convert the integer accumulator to a string and print each character
	printDigits(acc)
	z01.PrintRune('\n')
}

func main() {
	mul := func(acc int, cur int) int {
		return acc * cur
	}
	sum := func(acc int, cur int) int {
		return acc + cur
	}
	div := func(acc int, cur int) int {
		return acc / cur
	}
	sub := func(acc int, cur int) int {
		return acc - cur
	}

	as := []int{600, 2}
	ReduceInt(as, mul) // Output: 1000
	ReduceInt(as, sum) // Output: 502
	ReduceInt(as, div) // Output: 250
	ReduceInt(as, sub) // Output: 598
}
