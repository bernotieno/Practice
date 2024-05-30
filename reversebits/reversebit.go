package main

import "fmt"

func ReverseBits(oct byte) byte {
	var reversed byte = 0
	for i := 1; i <= 7; i++ {
		bit := (oct >> i) & 1

		reversed = reversed | (bit << (7 - i))
	}
	return reversed
}

func main() {
	var b byte = 0b00100110

	fmt.Printf("%08b", ReverseBits(b))
}
