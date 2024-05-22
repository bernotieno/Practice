// the code defines a Go program that finds the maximum value in a slice of integers
package main

import "fmt"

func Max(a []int) int {
	// Check if the input slice is empty
	if len(a) == 0 {
		fmt.Println(0)
	}

	// Initialize max variable with the first element of the slice
	max := a[0]

	for _, v := range a {
		// Compare each element with the current maximum
		if v > max {
			// If the current element is greater than the current maximum,
			// update the max variable to the value of the current element
			max = v
		}
	}

	// After iterating over all elements, return the maximum value found
	return max
}

func main() {
	a := []int{23, 123, 1, 11, 55, 93}
	max := Max(a)
	fmt.Println(max)
}
