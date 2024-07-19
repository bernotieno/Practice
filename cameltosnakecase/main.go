package main

import (
	"fmt"
)

// CamelToSnakeCase converts a camelCase string to snake_case
func CamelToSnakeCase(s string) string {
	if len(s) == 0 {
		return ""
	}

	var result string

	for i := 0; i < len(s); i++ {
		char := s[i]
		// Check if the current character is an uppercase letter
		if char >= 'A' && char <= 'Z' {
			// If previous character is a lowercase letter, add an underscore
			if i > 0 && s[i-1] >= 'a' && s[i-1] <= 'z' {
				result += "_"
			}
		}
		// Add the current character to the result
		result += string(char)
	}

	return result
}

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))       // Hello_World
	fmt.Println(CamelToSnakeCase("helloWorld"))       // hello_World
	fmt.Println(CamelToSnakeCase("camelCase"))        // camel_Case
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE")) // CAMELtoSnackCASE
	fmt.Println(CamelToSnakeCase("camelToSnakeCase")) // camel_To_Snake_Case
	fmt.Println(CamelToSnakeCase("hey2"))             // hey2
}
