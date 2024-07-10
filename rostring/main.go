package main

import (
	"fmt"
	"os"
)

func rostring(str string) string {
	var words []string
	var word string
	var result string

	// Split the input string into words
	for _, char := range str {
		if char != ' ' {
			word += string(char)
		} else if word != "" {
			words = append(words, word)
			word = ""
		}
	}
	if word != "" {
		words = append(words, word)
	}

	// Handle the case where there might be no words
	if len(words) == 0 {
		return ""
	}

	// Rotate the words
	words = append(words[1:], words[0])

	// Construct the result string without extra space
	for i, wor := range words {
		if i == len(words)-1 {
			result += wor
		} else {
			result += wor + " "
		}
	}
	return result
}

func main() {
	input := os.Args[1]
	fmt.Println(rostring(input))
}
