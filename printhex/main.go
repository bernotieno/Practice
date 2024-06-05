package main

import (
	"fmt"
	"os"
	"strconv"
)

func printhex(n int) string {
	hex := "0123456789abcdef"
	mod := []int{}
	result := ""

	for i := n; i > 0; i /= 16 {
		mod = append(mod, i%16)
	}

	for j := len(mod) - 1; j >= 0; j-- {
		index := mod[j]
		result += string(hex[index])
	}
	return result
}

func main() {
	if len(os.Args) != 2 {
		return
	}
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("ERROR")
	}
	fmt.Println(printhex(num))
}
