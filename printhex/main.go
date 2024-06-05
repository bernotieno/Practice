package main

import (
	"github.com/01-edu/z01"
)

// printhex converts an integer to its hexadecimal representation and prints it.
func printhex(n int) {
	// Define the hexadecimal characters.
	hex := "0123456789abcdef"
	output := ""

	// Slice to store the remainders.
	mod := []int{}

	// Convert the integer to hexadecimal by repeatedly dividing by 16
	// and storing the remainders.
	for i := n; i > 0; i /= 16 {
		mod = append(mod, i%16)
	}

	// Build the hexadecimal string by mapping the remainders to hex characters.
	for j := len(mod) - 1; j >= 0; j-- {
		ind := mod[j]
		output = output + string(hex[ind])
	}

	// Print the hexadecimal string using z01.PrintRune.
	for _, char := range output {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n') 
}

func main() {
	// Call printhex with the value 44, which should output '2c'.
	printhex(44)
}
