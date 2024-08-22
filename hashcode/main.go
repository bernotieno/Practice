package main

import (
	"fmt"
)

// HashCode takes a string and returns a new hashed string.
func HashCode(dec string) string {
	size := len(dec)
	hashed := make([]rune, size)

	for i, char := range dec {
		newChar := (int(char) + size) % 127
		if newChar < 33 || newChar > 126 {
			newChar += 33
		}
		hashed[i] = rune(newChar)
	}

	return string(hashed)
}

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BaC"))
	fmt.Println(HashCode("Hello World"))
}
