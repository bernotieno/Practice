// The ChunkSlice function is designed to divide a given slice of integers into smaller chunks of a specified size and print these chunks.

package main 

import (
	"fmt" 
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

	fmt.Println(chunks) 
}

func main() {
	ChunkSlice([]int{24, 23, 23, 2, 42, 23, 2, 432, 43, 2, 4, 2223, 42, 4, 243}, 3)
}
