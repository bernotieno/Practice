package main

import "fmt"

// Slice extracts a sub-slice from a given slice of strings based on the provided indices.
func Slice(a []string, nbrs ...int) []string {
	length := len(a)

	if len(nbrs) == 0 {
		return a
	}

	start := nbrs[0]
	end := length

	if len(nbrs) > 1 {
		end = nbrs[1]
	}

	// Handle negative indices
	if start < 0 {
		start = length + start
	}
	if end < 0 {
		end = length + end
	}

	// Ensure start and end are within bounds
	if start < 0 {
		// start = 0
		return []string{}
	}
	if end > length {
		// end = length
		return []string{}
	}

	// Return nil if start is greater than end or length
	if start >= length || start >= end {
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
