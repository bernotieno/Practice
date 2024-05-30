package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Atoi(s string) int {
	num := 0

	for _, v := range s {
		n := int(v - 48)

		if n < 0 || n > 9 {
			return 0
		}
		num = (num * 10) + n

	}
	return num
}

func main() {
	if len(os.Args) != 2 {
		return
	}
	input := Atoi(os.Args[1])
	var output1 []string
	var output2 []string

	if input == 0 {
		for _, c := range "ERROR: cannot convert to roman digit" {
			z01.PrintRune(c)
		}
		z01.PrintRune('\n')
		return
	}

	if input >= 4000 {
		for _, b := range "ERROR: cannot convert to roman digit" {
			z01.PrintRune(b)
		}
		z01.PrintRune('\n')
		return
	}

	for input > 0 {
		if input >= 1000 {
			output1 = append(output1, "M")
			output2 = append(output2, "M")
			input -= 1000
		} else if input >= 900 {
			output1 = append(output1, "CM")
			output2 = append(output2, "(M-C)")
			input -= 900
		} else if input >= 500 {
			output1 = append(output1, "D")
			output2 = append(output2, "D")
			input -= 500
		} else if input >= 400 {
			output1 = append(output1, "CD")
			output2 = append(output2, "(D-C)")
			input -= 400
		} else if input >= 100 {
			output1 = append(output1, "C")
			output2 = append(output2, "C")
			input -= 100
		} else if input >= 90 {
			output1 = append(output1, "XC")
			output2 = append(output2, "(C-X)")
			input -= 90
		} else if input >= 50 {
			output1 = append(output1, "L")
			output2 = append(output2, "L")
			input -= 50
		} else if input >= 40 {
			output1 = append(output1, "XL")
			output2 = append(output2, "(L-X)")
			input -= 40
		} else if input >= 10 {
			output1 = append(output1, "X")
			output2 = append(output2, "X")
			input -= 10
		} else if input >= 9 {
			output1 = append(output1, "IX")
			output2 = append(output2, "(X-I)")
			input -= 9
		} else if input >= 5 {
			output1 = append(output1, "V")
			output2 = append(output2, "V")
			input -= 5
		} else if input >= 4 {
			output1 = append(output1, "IV")
			output2 = append(output2, "(V-I)")
			input -= 4
		} else if input >= 1 {
			output1 = append(output1, "I")
			output2 = append(output2, "I")
			input -= 1
		}
	}

	for i, slice := range output2 {
		for _, d := range slice {
			z01.PrintRune(d)
		}
		if i != len(output2)-1 {
			z01.PrintRune('+')
		}
	}
	z01.PrintRune('\n')

	for _, slice2 := range output1 {
		for _, e := range slice2 {
			z01.PrintRune(e)
		}
	}
	z01.PrintRune('\n')
}
