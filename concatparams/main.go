package main

import "fmt"

func ConcatParams(args []string) string {
	if len(args) == 0 {
		return ""
	}
	var result string
	for i, arg := range args {
		result += arg
		if i < len(args)-1 {
			result += "\n"
		}
	}
	return result
}

func main() {
	test := []string{"Hello", "how", "are", "you?"}
	fmt.Println(ConcatParams(test))
}
