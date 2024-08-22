package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("options: -abcdefghijklmnopqrstuvwxyz")
		return
	}
	args := os.Args[1:]
	option(args)
}

func option(str []string) {
	var optionsInt uint32 = 0

	for _, arg := range str {
		if len(arg) > 0 && arg[0] == '-' {
			if len(arg) > 1 && contains(arg) {
				fmt.Println("options: -abcdefghijklmnopqrstuvwxyz")
				return
			}
			if arg[0] == '-' && len(arg) == 1 {
				fmt.Println("Invalid Option")
				return
			}
			for _, ch := range arg[1:] {
				if ch >= 'a' && ch <= 'z' {
					bitposition := ch - 'a'
					optionsInt |= (1 << bitposition)
				} else {
					fmt.Println("Invalid Option")
					return
				}
			}
		} else {
			fmt.Println("Invalid Option")
			return
		}
	}

	bitString := intToBinaryString(optionsInt, 32)

	for i := 0; i < 32; i += 8 {
		if (i + 8) < 32 {
			fmt.Print(bitString[i:i+8] + " ")
		} else {
			fmt.Print(bitString[i : i+8])
		}
	}
	fmt.Println()
}

func intToBinaryString(n uint32, bitSize int) string {
	binary := make([]byte, bitSize)
	for i := 0; i < bitSize; i++ {
		if n&(1<<(bitSize-i-1)) != 0 {
			binary[i] = '1'
		} else {
			binary[i] = '0'
		}
	}
	return string(binary)
}

func contains(s string) bool {
	return s[:2] == "-h"
}
