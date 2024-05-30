package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	str1 := os.Args[1]
	str2 := os.Args[2]
	seen := make(map[rune]bool)
	exist := make(map[rune]bool)
	inter := ""

	for _, char := range str2 {
		exist[char] = true
	}

	for _, char := range str1 {
		if exist[char] && !seen[char] {
			seen[char] = true
			inter += string(char)
		}
	}

	for _, c := range inter {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
