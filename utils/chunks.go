package functions // Declares the main package

import (
	"fmt" // Imports the fmt package for formatted I/O
)

func ChunkSlice(slice []int, chunkSize int) {
	var chunks [][]int // Declares a variable to hold the chunks

	// Iterates over the slice in chunks of specified size
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize // Calculates the end index of the current chunk

		if end > len(slice) { // Checks if the end index exceeds the length of the slice
			end = len(slice) // Adjusts the end index to the length of the slice
		}

		chunks = append(chunks, slice[i:end]) // Appends the current chunk to the chunks slice
	}

	fmt.Println(chunks) // Returns the slice of chunks
}
