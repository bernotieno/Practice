package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	count := 0

	for i := 1; i < len(os.Args); i++ {
		count++
	}

	// Print the count
	z01.PrintRune(rune(count + '0'))
	z01.PrintRune('\n')
}
