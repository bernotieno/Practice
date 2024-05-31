package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Print(str string) {
	for _, char := range str {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}

func main() {
	if len(os.Args) != 3 {
		return
	}
	args := os.Args[1:]
	union := make(map[rune]bool)
	unionstr := ""

	for _, c := range args[0] + args[1] {
		if !union[c] {
			unionstr += string(c)
			union[c] = true
		}
	}

	Print(unionstr)
}
