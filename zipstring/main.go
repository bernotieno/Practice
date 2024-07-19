package main

import (
	"strconv"
)

func ZipString(s string) string {
	if len(s) == 0 {
		return ""
	}

	var result string
	count := 1

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			result += strconv.Itoa(count) + string(s[i-1])
			count = 1
		}
	}
	// Append the last group
	result += strconv.Itoa(count) + string(s[len(s)-1])

	return result
}
