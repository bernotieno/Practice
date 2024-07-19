package main

import "fmt"

// CanJump determines if you can reach the last index of the array
func CanJump(steps []uint) bool {
	n := len(steps)

	// Edge case: if the array is empty, return false
	if n == 0 {
		return false
	}

	// Edge case: if there's only one element, you're already at the last position
	if n == 1 {
		return true
	}

	// Initialize the maximum reachable index
	maxReachable := 0

	for i := 0; i < n; i++ {
		// If the current index is beyond the maximum reachable index, return false
		if i > maxReachable {
			return false
		}

		// Update the maximum reachable index
		if i+int(steps[i]) > maxReachable {
			maxReachable = i + int(steps[i])
		}

		// If the maximum reachable index is greater than or equal to the last index, return true
		if maxReachable >= n-1 {
			return true
		}
	}

	// If we have gone through all positions and didn't return true, return false
	return false
}

func main() {
	input1 := []uint{2, 3, 1, 1, 4}
	fmt.Println(CanJump(input1))

	input2 := []uint{3, 2, 1, 0, 4}
	fmt.Println(CanJump(input2))

	input3 := []uint{0}
	fmt.Println(CanJump(input3))
}
