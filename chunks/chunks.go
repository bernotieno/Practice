package main

import "fmt"

// The ChunkSlice function is designed to divide a given slice of integers into smaller chunks of a specified size and print these chunks.
func ChunkSlice(slice []int, chunkSize int) {
	if chunkSize <= 0 {
		fmt.Println()
		return
	}
	if len(slice) == 0 {
		return
	}
	var chunks [][]int

	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize
		if end > len(slice) {
			end = len(slice)
		}

		chunks = append(chunks, slice[i:end])
	}

	fmt.Println(chunks)
}

func main() {
	ChunkSlice([]int{24, 23, 23, 2, 42, 23, 2, 432, 43, 2, 4, 2223, 42, 4, 243}, 3)
}
