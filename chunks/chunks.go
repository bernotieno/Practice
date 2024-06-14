// The ChunkSlice function is designed to divide a given slice of integers into smaller chunks of a specified size and print these chunks.

package main

import (
	"github.com/01-edu/z01"
)

func ChunkSlice(slice []int, chunkSize int) {
	var chunks [][]int

	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize // Calculates the end index of the current chunk

		if end > len(slice) { // Checks if the end index exceeds the length of the slice
			end = len(slice) // Adjusts the end index to the length of the slice
		}

		chunks = append(chunks, slice[i:end]) // Appends the current chunk to the chunks slice
	}

	Printit(chunks)
}

func Printit(chunkies [][]int) {
	z01.PrintRune('[')
	for c, chunk := range chunkies {
		z01.PrintRune('[')
		for i, num := range chunk {
			if i > 0 {
				z01.PrintRune(' ')
			}
			numString := Itoa(num)
			for _, r := range numString {
				z01.PrintRune(r)
			}
		}
		z01.PrintRune(']')
		if c < len(chunkies)-1 {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune(']')
	z01.PrintRune('\n')
}

func Itoa(nb int) string {
	if nb == 0 {
		return "0"
	}
	isNeg := false
	if nb < 0 {
		isNeg = true
		nb = -nb
	}
	var result []rune
	for nb > 0 {
		digit := nb % 10
		result = append([]rune{rune(digit + '0')}, result...)
		nb /= 10
	}

	if isNeg {
		result = append([]rune{'-'}, result...)
	}
	return string(result)
}

func main() {
	ChunkSlice([]int{24, 23, 23, 2, 42, 23, 2, 432, 43, 2, 4, 2223, 42, 4, 243}, 3)
}
