package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
		return
	}

	stInt := os.Args[1]
	var number int
	var outPut []string
	for i, v := range stInt {
		if i == 0 && v == '-' {
			return
		}
		n := int(v - 48)
		number = ((number * 10) + n)
	}

	for i := 1; i <= 9; i++ {
		newStr := string(rune(i + 48))
		res := (i * number)
		resStr := ""

		for res > 0 {
			digit := res % 10
			resStr = string(rune(digit+'0')) + resStr
			res /= 10
		}

		outPut = append(outPut, newStr+" "+"x"+" "+stInt+" "+"="+" "+resStr)

	}
	for i := 0; i < len(outPut); i++ {
		for _, v := range outPut[i] {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
	}
}
