package functions

import "github.com/01-edu/z01"

func FirstRune(str string) {
	char := str[0]

	z01.PrintRune(rune(char))
	z01.PrintRune('\n')
}
