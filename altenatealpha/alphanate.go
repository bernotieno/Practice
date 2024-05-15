package main

import (
	"github.com/01-edu/z01"
)

func main() {
	for i := 'A'; i <= 'Z'; i++ {
		if i%2 == 0 {
			z01.PrintRune(i)
		} else {
			z01.PrintRune(i + 32) // Convert uppercase to lowercase by adding 32 to ASCII value
		}
	}
	z01.PrintRune('\n') // Add a newline at the end
}
