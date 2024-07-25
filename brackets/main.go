package main

import (
	"fmt"
	"os"
)

func correctBracket(str string) string {
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
				return "Error"
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) == 0 {
		return "OK"
	}
	return "Error"
}

func main() {
	if len(os.Args) <= 1 {
		return
	}

	for _, args := range os.Args[1:] {
		fmt.Println(correctBracket(args))
	}
}
