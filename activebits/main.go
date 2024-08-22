package main

import "fmt"

func ActiveBits(n int) int {
	count := 0

	for i := 0; i < 32; i++ {
		bit := (n >> i) & 1
		if bit == 1 {
			count++
		}
	}
	return count
}

func main() {
	num := 7
	fmt.Println(ActiveBits(num)) // Output: 3, since 7 in binary is 111

	num = 15
	fmt.Println(ActiveBits(num)) // Output: 4, since 15 in binary is 1111
}
