package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 2 {
		n, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("ERROR")
			return
		}
		res := printHex(n, "0123456789abcdef")
		fmt.Println(res)
	}
}

func printHex(num int, base string) string {
	result := ""

	for num > 0 {
		rem := num % len(base)
		result = string(base[rem]) + result
		num /= len(base)
	}
	return result
}
