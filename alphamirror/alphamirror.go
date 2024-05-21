// The alphamirror exercise involves writing a function that mirrors the alphabetic characters in a string.
//  For example, 'a' becomes 'z', 'b' becomes 'y', and so on.
//  Non-alphabetic characters remain unchanged.

package main

import (
	"github.com/01-edu/z01"
)

// Function to perform alphamirror encryption
func Alphamirror(s string) {
	// Initialize a variable to store the encrypted string
	var result string

	// Iterate through each character in the input string
	for _, char := range s {
		// Check if the character is an uppercase letter
		if char >= 'A' && char <= 'Z' {
			// Calculate the mirror character by subtracting the ASCII value from 'Z'
			mirror := 'Z' - (char - 'A')
			// Append the mirror character to the result string
			result += string(mirror)
		} else if char >= 'a' && char <= 'z' { // Check if the character is a lowercase letter
			// Calculate the mirror character by subtracting the ASCII value from 'z'
			mirror := 'z' - (char - 'a')
			// Append the mirror character to the result string
			result += string(mirror)
		} else { // If the character is not a letter, append it unchanged
			result += string(char)
		}
	}

	for _, char := range result {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}

func main() {
	Alphamirror("a")
}
