package main

import (
	"fmt"
)

func WeAreUnique(str1, str2 string) int {
	if str1 == "" && str2 == "" {
		return -1
	}

	// Create maps to store unique characters
	set1 := make(map[rune]bool)
	set2 := make(map[rune]bool)

	// Populate the sets with unique characters
	for _, ch := range str1 {
		set1[ch] = true
	}
	for _, ch := range str2 {
		set2[ch] = true
	}

	// Create a map to track unique characters across both strings
	unique := make(map[rune]bool)

	// Add unique characters from set1 that are not in set2
	for ch := range set1 {
		if !set2[ch] {
			unique[ch] = true
		}
	}

	// Add unique characters from set2 that are not in set1
	for ch := range set2 {
		if !set1[ch] {
			unique[ch] = true
		}
	}

	// Return the count of unique characters
	return len(unique)
}

func main() {
	fmt.Println(WeAreUnique("foo", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "def"))
}
