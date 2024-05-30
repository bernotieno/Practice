package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Check if the number of command-line arguments is exactly 4
	if len(os.Args) != 4 {
		return // Exit the function if the arguments are not as expected
	}
	args := os.Args[1:] // Exclude the program name from the arguments
	var words string

	// Loop through each character in the first argument
	for _, arg := range args[0] {
		// If the character matches the second argument
		if string(arg) == args[1] {
			words += args[2] // Replace the character with the third argument
		} else {
			words += string(arg) // Keep the character unchanged
		}
	}
	// Print the modified string
	for _, c := range words {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n') // Print a newline character after the modified string
}
