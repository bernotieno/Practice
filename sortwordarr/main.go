package main

import (
	"fmt"
)

// SortWordArr sorts a slice of strings in ascending order using a simple
// bubble sort algorithm. The function modifies the input slice in place.
func SortWordArr(a []string) {
	for i := 0; i <= len(a)-1; i++ {
		for j := i + 1; j <= len(a)-1; j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}

func main() {
	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	SortWordArr(result)

	fmt.Println(result)
}
