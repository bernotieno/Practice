package main

func AddFront(s string, slice []string) []string {
	// Check if the string s is empty.
	if s == "" {
		return slice
	}

	// Create a new slice with s as the first element and append the original slice.
	newSlice := append([]string{s}, slice...)
	return newSlice
}
