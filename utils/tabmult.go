package functions

import (
	"os"

	"github.com/01-edu/z01"
)

// stringToRuneSlice converts a string to a slice of runes.
func stringToRuneSlice(s string) []rune {
	runes := make([]rune, len(s))
	for i, r := range s {
		runes[i] = r
	}
	return runes
}

// runeSliceToInt converts a slice of runes to an integer.
func runeSliceToInt(runes []rune) int {
	num := 0
	for _, r := range runes {
		num = num*10 + int(r-'0')
	}
	return num
}

// intToString converts an integer to a string.
func intToString(num int) string {
	if num == 0 {
		return "0"
	}
	var result []rune
	isNegative := false
	if num < 0 {
		isNegative = true
		num = -num
	}
	for num > 0 {
		digit := '0' + rune(num%10)
		result = append(result, digit)
		num /= 10
	}
	if isNegative {
		result = append(result, '-')
	}
	// Reverse the rune slice
	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-1-i] = result[len(result)-1-i], result[i]
	}
	return string(result)
}

// Tabmult prints the multiplication table for a given number.
func Tabmult() {
	// Check if exactly one argument is provided
	if len(os.Args) != 2 {
		return
	}

	// Convert the argument to a rune slice
	argRunes := stringToRuneSlice(os.Args[1])

	// Convert the rune slice to an integer
	num := runeSliceToInt(argRunes)

	// Print the multiplication table for the given number
	for i := 1; i <= 9; i++ {
		// Calculate the result of the multiplication
		result := i * num

		// Print the formatted multiplication table entry
		z01.PrintRune(rune('0' + i)) // Print the multiplier
		z01.PrintRune(' ')
		z01.PrintRune('x')
		z01.PrintRune(' ')
		for _, r := range argRunes {
			z01.PrintRune(r)
		}
		z01.PrintRune(' ')
		z01.PrintRune('=')
		z01.PrintRune(' ')
		for _, r := range stringToRuneSlice(intToString(result)) {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
