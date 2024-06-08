package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	str := os.Args[1]
	result := ""
	if str == "" {
		z01.PrintRune('\n')
		return
	}
	for _, char := range str {
		if char >= 'A' && char <= 'Z' {
			index := int(char - 'A')
			for i := 0; i <= index; i++ {
				result += string(char)
			}
		} else if char >= 'a' && char <= 'z' {
			index := int(char - 'a')
			for i := 0; i <= index; i++ {
				result += string(char)
			}
		} else {
			result += string(char)
		}
	}
	Printn(result)
}

func Printn(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
