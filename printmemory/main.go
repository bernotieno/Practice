package main

import "github.com/01-edu/z01"

func printstr(str string) {
	for _, char := range str {
		z01.PrintRune(char)
	}
}

func hex(str string) string {
	hex := "0123456789abcdef"
	var result string

	if str == "" {
		return "00"
	}

	for _, char := range str {
		for char > 0 {
			remainder := int(char) % 16

			result = string(hex[remainder]) + result
			char /= 16
		}
	}
	return result
}

func PrintMemory(arr [10]byte) {
	for i, b := range arr {
		// byte 0
		if b == '\x00' {
			printstr("00")
		}

		// fmt.Printf("%02x", b)
		printstr(hex(string(b)))

		if (i+1)%4 == 0 {
			z01.PrintRune('\n')
		} else if i != len(arr)-1 {
			z01.PrintRune(' ')
		}

	}
	printstr("\n")

	for _, b := range arr {
		if b >= 32 && b <= 126 {
			printstr(string(b))
		} else {
			printstr(".")
		}
	}
	printstr("\n")
}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}
