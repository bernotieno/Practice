package functions

import "fmt"

// Function to reduce a slice of integers to a single integer by summing them up
func Reduceint(nums []int) {
	// Initialize a variable to hold the sum
	sum := 0

	// Iterate through the slice and add each element to the sum
	for _, num := range nums {
		sum += num
	}

	// Return the sum as the reduced value
	fmt.Println(sum)
}
