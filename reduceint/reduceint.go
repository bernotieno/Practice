// Overall, this program showcases how to implement a generic reduction function in Go,
// allowing for different reduction operations to be applied to a slice of integers,
// with the flexibility to print the result in a custom format using the z01 package.

package main

import "github.com/01-edu/z01"

func printInt(n int) {
	// Handle zero case separately
	if n == 0 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}

	// Convert integer to slice of runes (characters)
	var digits []rune
	for n > 0 {
		digit := n % 10
		digits = append([]rune{rune(digit + '0')}, digits...) // Prepend digit to the slice
		n /= 10
	}

	// Print each rune
	for _, r := range digits {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
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
	printInt(acc)
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
