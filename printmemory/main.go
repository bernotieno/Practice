package main

import "github.com/01-edu/z01"

func printstr(str string) {
	for _, char := range str {
		z01.PrintRune(char)
	}
}

func hex(b byte) string {
	hexDigits := "0123456789abcdef"
	return string(hexDigits[b>>4]) + string(hexDigits[b&0x0f])
}

func PrintMemory(arr [10]byte) {
	for i, b := range arr {
		// Print each byte as two hexadecimal digits
		printstr(hex(b))

		if (i+1)%4 == 0 || i == len(arr)-1 {
			z01.PrintRune('\n')
		} else {
			z01.PrintRune(' ')
		}
	}

	for _, b := range arr {
		if b > 32 && b <= 126 {
			printstr(string(b))
		} else {
			printstr(".")
		}
	}

	z01.PrintRune('\n')
}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}
