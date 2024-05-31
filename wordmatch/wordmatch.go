package main

import (
	"fmt"
	"os"
)

func containWord(first, second string) bool {
	j := 0
	for i := 0; i < len(second) && j < len(first); i++ {
		if first[j] == second[i] {
			j++
		}
	}
	return j == len(first)
}

func WordMatch(first, second string) string {
	if containWord(first, second) {
		return first
	}
	return ""
}

func main() {
	if len(os.Args) != 3 {
		return
	}
	str1 := os.Args[1]
	str2 := os.Args[2]
	fmt.Println(WordMatch(str1, str2))
}
