package main

import (
	"fmt"
)

func AtoiBase(s, base string) int {
	if !isValidBase(base) {
		return 0
	}
	baseLen := len(base)
	index := 0
	res := 0
	for _, char := range s {
		for pos, baze := range base {
			if char == baze {
				index = pos
			}
		}
		res = res*baseLen + index
	}
	return res
}

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}

	baseMap := make(map[rune]bool)

	for _, char := range base {
		if char == '+' || char == '-' {
			return false
		}
		if baseMap[char] {
			return false
		}
		baseMap[char] = true
	}
	return true
}

func main() {
	fmt.Println(AtoiBase("125", "0123456789"))
	fmt.Println(AtoiBase("1111101", "01"))
	fmt.Println(AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(AtoiBase("uoi", "choumi"))
	fmt.Println(AtoiBase("bbbbbab", "-ab"))
}
