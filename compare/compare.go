// The Compare function is a simple Go function that compares two strings and prints a message based on their comparison

package main

import "fmt"

func Compare(a, b string) {
	// Check if strings a and b are equal
	if len(a) == len(b) {
		fmt.Println("equal strings")
	} else if len(a) > len(b) {
		fmt.Println("first longer than second")
	} else if len(a) < len(b) {
		fmt.Println("second is longer than first")
	}
}

func main() {
	Compare("hello", "beautifully")
}
