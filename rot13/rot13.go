package main

import (
	"github.com/01-edu/z01"
)

func Rot13(s string) {
	var result string
	for _, char := range s {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			// Apply ROT13 transformation for letters
			offset := 'a'
			if char >= 'A' && char <= 'Z' {
				offset = 'A'
			}
			result += string((char-offset+13)%26 + offset)
		} else {
			// If the character is not a letter, leave it unchanged
			result += string(char)
		}
	}
	for _, char := range result {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}

func main() {
	Rot13("a")
}
