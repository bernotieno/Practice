package functions

import "fmt"

func Compare(a, b string) {
	// Check if strings a and b are equal
	if a == b {
		fmt.Println("equal strings")
	} else if a > b {
		// Return 1 if string a is greater than string b
		fmt.Println("1st longer than second")
	} else {
		// Return -1 if string a is less than string b
		fmt.Println("1st smaller than second")
	}
}
