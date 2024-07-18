package main

import (
	"fmt"
	"os"
)

func split(str, sep string) []string {
	var result []string
	start := 0
	for i := 0; i+len(sep) <= len(str); i++ {
		if str[i:i+len(sep)] == sep {
			if start < i {
				result = append(result, str[start:i])
			}
			start = i + len(sep)
		}
	}
	if start < len(str) {
		result = append(result, str[start:])
	}
	return result
}

func rot(slice []string) []string {
	rotate := make([]string, len(slice))
	for range slice {
		rotate = append(slice[1:], slice[0])
	}
	return rotate
}

func join(slice []string) string {
	var final string
	for i, word := range slice {
		final += word
		if i < len(slice)-1 {
			final += " "
		}
	}
	return final
}

func main() {
	if len(os.Args) != 2 || os.Args[1] == "" {
		return
	}
	args := os.Args[1]

	splitted := split(args, " ")
	rotated := rot(splitted)
	joined := join(rotated)
	fmt.Println(joined)
}
