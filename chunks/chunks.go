package main

import "fmt"

// The ChunkSlice function is designed to divide a given slice of integers into smaller chunks of a specified size and print these chunks.
func ChunkSlice(slice []int, chunkSize int) {
	if chunkSize <= 0 {
		fmt.Println()
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
	ChunkSlice([]int{}, 10)
	ChunkSlice([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	ChunkSlice([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	ChunkSlice([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	ChunkSlice([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}
