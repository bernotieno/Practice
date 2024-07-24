package main

import "fmt"

func WordFlip(str string) string {
	if str == "" {
		return "Invalid Output\n"
	}
	split := split(str, " ")
	rev := reverseSlice(split)
	joined := joinWithSpaces(rev)
	if joined == " " {
		return ""
	}
	return joined + "\n"
}

func split(str, sep string) []string {
	var result []string
	start := 0
	for i := 0; i+len(sep) <= len(str); i++ {
		if str[i:i+len(sep)] == sep {
			if start < i {
				result = append(result, str[start:i])
			}
			start = i + len(sep)
		}
	}
	if start < len(str) {
		result = append(result, str[start:])
	}
	return result
}

// reverseSlice takes a slice of strings and returns a new slice of strings with the elements in reverse order.
func reverseSlice(slice []string) []string {
	reversed := make([]string, len(slice))
	for i := range slice {
		reversed[i] = slice[len(slice)-1-i]
	}
	return reversed
}

// joinWithSpaces takes a slice of strings and returns a single string with the elements joined by spaces.
func joinWithSpaces(slice []string) string {
	var result string
	for i, word := range slice {
		result += word
		if i < len(slice)-1 {
			result += " "
		}
	}
	return result
}

func main() {
	fmt.Print(WordFlip("First second last"))
	fmt.Print(WordFlip(""))
	fmt.Print(WordFlip("     "))
	fmt.Print(WordFlip(" hello  all  of  you! "))
}
