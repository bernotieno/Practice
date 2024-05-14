package main

import (
	"fmt"
)

func main() {
	// Sample slice of strings
	words := []string{"apple", "banana", "orange", "grape", "pineapple"}

	// Word to match
	target := "banana"

	// Call the function to check if the target word exists in the slice
	if wordMatch(words, target) {
		fmt.Printf("'%s' exists in the slice.\n", target)
	} else {
		fmt.Printf("'%s' does not exist in the slice.\n", target)
	}
}

// Function to check if a word exists in a slice of strings
func wordMatch(words []string, target string) bool {
	// Iterate through the slice of strings
	for _, word := range words {
		// Check if the current word matches the target word
		if word == target {
			return true // Return true if a match is found
		}
	}
	return false // Return false if no match is found
}
