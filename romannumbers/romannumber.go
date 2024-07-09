package main

import (
	"fmt"
	"os"
	"strconv"
)

func toRoman(nb int) (string, string) {
	calculationSymbol := []string{"M", "(M-C)", "D", "(D-C)", "C", "(C-X)", "L", "(L-X)", "X", "(X-I)", "V", "(V-I)", "I"}
	value := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbol := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	roman := ""
	calculation := ""

	for i := 0; i < len(value); i++ {
		for nb >= value[i] {
			nb -= value[i]
			roman += symbol[i]
			if calculation != "" {
				calculation += "+"
			}
			calculation += calculationSymbol[i]
		}
	}
	return calculation, roman
}

func main() {
	input, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("ERROR: cannot convert to roman digit")
		return
	}
	if input <= 0 || input >= 4000 {
		fmt.Println("ERROR: cannot convert to roman digit")
		return
	}
	rom, cal := toRoman(input)
	fmt.Println(rom)
	fmt.Println(cal)
}
