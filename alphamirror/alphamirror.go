package main

import (
	"os"

	"github.com/01-edu/z01"
)

// The alphamirror exercise involves writing a function that mirrors the alphabetic characters in a string.
func Alphamirror(s string) {
	var result string

	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			mirror := 'Z' - (char - 'A')
			result += string(mirror)
		} else if char >= 'a' && char <= 'z' {
			mirror := 'z' - (char - 'a')
			result += string(mirror)
		} else {
			result += string(char)
		}
	}

	for _, char := range result {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}

func main() {
	if len(os.Args) != 2 {
		return
	}
	input := os.Args[1]
	Alphamirror(input)
}
