package main

import "fmt"

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
	fmt.Println(WordMatch("hello", "djdhndedkdldjdldhos"))
}
