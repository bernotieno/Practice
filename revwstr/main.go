package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 || os.Args[1] == "" {
		return
	}

	args := os.Args[1]
	words := splitIntoWords(args)
	revWords := reverseSlice(words)
	revStr := joinWithSpaces(revWords)
	printString(revStr)
}

// splitIntoWords function, which takes a string s and returns a slice of strings.
func splitIntoWords(s string) []string {
	var words []string
	start := 0
	for i, c := range s {
		if c == ' ' {
			words = append(words, s[start:i])
			start = i + 1
		}
	}
	words = append(words, s[start:])
	return words
}

// reverseSlice takes a slice of strings and returns a new slice of strings with the elements in reverse order.
func reverseSlice(slice []string) []string {
	reversed := make([]string, len(slice))
	for i := range slice {
		reversed[i] = slice[len(slice)-1-i]
	}
	return reversed
}

// joinWithSpaces takes a slice of strings and returns a single string with the elements joined by spaces.
func joinWithSpaces(slice []string) string {
	var result string
	for i, word := range slice {
		result += word
		if i < len(slice)-1 {
			result += " "
		}
	}
	return result
}

// printString function, which takes a string s and prints each character as a rune.
func printString(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
