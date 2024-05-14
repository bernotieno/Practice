package main

import "fmt"

func main() {
	// Sample slice of integers
	numbers := []int{5, 10, 15, 20}

	// Call the function to reduce the slice to a single integer
	result := reduce(numbers)

	// Print the result
	fmt.Println("The reduced value is:", result)
}

// Function to reduce a slice of integers to a single integer by summing them up
func reduce(nums []int) int {
	// Initialize a variable to hold the sum
	sum := 0

	// Iterate through the slice and add each element to the sum
	for _, num := range nums {
		sum += num
	}

	// Return the sum as the reduced value
	return sum
}
