package main

import "fmt"

func PrintIfNot(str string) string {
	var result string

	if len(str) < 3 || str == "" {
		result += "G\n"
	} else {
		result += "Invalid Input\n"
	}
	return result
}

func main() {
	fmt.Print(PrintIfNot("abcdefz"))
	fmt.Print(PrintIfNot("abc"))
	fmt.Print(PrintIfNot(""))
	fmt.Print(PrintIfNot("14"))
}
