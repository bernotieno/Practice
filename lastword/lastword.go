package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	input := os.Args[1]
	if input == "" {
		return
	}

	sentence := []rune(input)
	var lastWord string

	// iterate over input string in reverse to find last word
	for i := len(sentence) - 1; i >= 0; i-- {
		// if space is encountered and word isn't empty == found last word
		if sentence[i] == ' ' {
			if len(lastWord) > 0 {
				break
			}
		} else {
			lastWord = string(sentence[i]) + lastWord // to print correctly
			// if you use word += sentence[i] // have to print in reverse
		}
	}
	for _, char := range lastWord {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
