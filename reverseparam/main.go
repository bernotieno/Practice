package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Loop through the arguments in reverse order
	for i := len(os.Args) - 1; i > 0; i-- {
		// Print each argument
		for _, char := range os.Args[i] {
			z01.PrintRune(char)
		}
		z01.PrintRune(' ') // Add a space between arguments
	}
	z01.PrintRune('\n') // Add a newline after printing all arguments
}
