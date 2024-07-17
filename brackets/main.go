package main

import (
	"fmt"
	"os"
)

func correctBracket(str string) bool {
	stack := []rune{}

	validB := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	for _, char := range str {
		switch char {
		case '(', '[', '{':
			stack = append(stack, char)
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != validB[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func main() {
	if len(os.Args) <= 1 {
		return
	}

	args := os.Args[1:]
	for _, words := range args {
		if correctBracket(words) || words == "" {
			fmt.Println("OK")
		} else {
			fmt.Println("ERROR")
		}
	}
}
