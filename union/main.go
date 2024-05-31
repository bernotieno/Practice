package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	args := os.Args[1:]

	union := make(map[rune]bool)
	unionStr := ""

	for _, c := range args[0] + args[1] {
		if !union[c] { // checks if the letter is not in the union map yet
			unionStr += string(c)
			union[c] = true // ones the letter is added the value becomes true
		}
	}

	for _, c := range unionStr {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
