package main

import (
	"fmt"
	"os"
)

func piglatin(s string) string {
	result := ""
	for i, char := range s {
		if i == 0 && checkVowel(char) {
			result = s + "ay"
			return result
		} else if i != 0 && checkVowel(char) {
			before := s[:i]
			after := s[i:]
			result = after + before + "ay"
			return result
		}
	}
	result = "No vowels"
	return result
}

// function to check for vowels
func checkVowel(s rune) bool {
	return s == 'a' || s == 'e' || s == 'i' || s == 'o' || s == 'u' || s == 'A' || s == 'E' || s == 'I' || s == 'O' || s == 'U'
}

func main() {
	if len(os.Args) != 2 {
		return
	}
	args := os.Args[1]
	fmt.Println(piglatin(args))
}
