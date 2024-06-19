package main

import "github.com/01-edu/z01"

func FirstRune(s string) rune {
	runes := []rune(s)
	return runes[0]
}

func main() {
	z01.PrintRune(FirstRune("♡Hello!"))
	z01.PrintRune('\n')
}
