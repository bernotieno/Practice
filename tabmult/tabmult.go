package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Check if the number of command-line arguments is exactly 2
	if len(os.Args) != 2 {
		z01.PrintRune('\n') // Print a newline and exit if arguments are incorrect
		return
	}

	stInt := os.Args[1] // Get the argument representing an integer
	var number int      // Initialize a variable to store the integer value
	var outPut []string // Initialize a slice to store the output strings

	// Iterate over each character in the input string
	for i, v := range stInt {
		// Skip the first character if it is a '-' (negative sign)
		if i == 0 && v == '-' {
			return
		}
		// Convert the character to its integer value
		n := int(v - 48)
		// Build the integer by accumulating digits
		number = ((number * 10) + n)
	}

	// Loop from 1 to 9 (inclusive)
	for i := 1; i <= 9; i++ {
		// Create a new string representing the current multiplier
		newStr := string(rune(i + 48))
		// Calculate the result of multiplication
		res := (i * number)
		resStr := "" // Initialize a string to store the result as a string

		// Convert the result to a string by extracting digits
		for res > 0 {
			digit := res % 10                         // Extract the last digit
			resStr = string(rune(digit+'0')) + resStr // Add the digit to the result string
			res /= 10                                 // Move to the next digit
		}

		// Build the output string and append it to the output slice
		outPut = append(outPut, newStr+" "+"x"+" "+stInt+" "+"="+" "+resStr)
	}

	// Print each string in the output slice
	for i := 0; i < len(outPut); i++ {
		for _, v := range outPut[i] {
			z01.PrintRune(v) // Print each character of the string
		}
		z01.PrintRune('\n') // Print a newline after each output string
	}
}
