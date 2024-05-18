package functions

import (
	"os"

	"github.com/01-edu/z01"
)

func Firstparam() {
	args := os.Args[1]
	for _, char := range args {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
