package main

import "fmt"

// Slice extracts a sub-slice from a given slice of strings based on the provided indices.
func Slice(a []string, nbrs ...int) []string {
	if len(nbrs) == 0 {
		return nil
	}

	// Handle start and end indices
	start, end := 0, len(a)

	if len(nbrs) > 0 {
		start = nbrs[0]
	}

	if len(nbrs) > 1 {
		end = nbrs[1]
	}

	// Convert negative indices to positive equivalents
	if start < 0 {
		start = len(a) + start
	}

	if end < 0 {
		end = len(a) + end
	}

	// Ensure indices are within bounds
	if start < 0 || start > len(a) || end < 0 || end > len(a) || start > end {
		return nil
	}

	return a[start:end]
}

func main() {
	a := []string{"coding", "algorithm", "ascii", "package", "golang"}
	fmt.Printf("%#v\n", Slice(a, 1))
	fmt.Printf("%#v\n", Slice(a, 2, 4))
	fmt.Printf("%#v\n", Slice(a, -3))
	fmt.Printf("%#v\n", Slice(a, -2, -1))
	fmt.Printf("%#v\n", Slice(a, 2, 0))
}
