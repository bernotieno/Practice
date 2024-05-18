package functions

import "fmt"

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
	fmt.Println(result)
}
