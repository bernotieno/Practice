package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	args := os.Args[1]
	final := repeatAlpha(args)
	printer(final)
}

func repeatAlpha(str string) string {
	finalStr := ""

	for _, char := range str {
		if char >= 'a' && char <= 'z' {
			index := int(char - 'a')
			for i := 0; i <= index; i++ {
				finalStr += string(char)
			}
		} else if char >= 'A' && char <= 'Z' {
			index2 := int(char - 'A')
			for j := 0; j <= index2; j++ {
				finalStr += string(char)
			}
		} else {
			finalStr += string(char)
		}
	}
	return finalStr
}

func printer(str string) {
	for _, ch := range str {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
