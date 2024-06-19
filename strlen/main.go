package main

import "fmt"

func Strlen(str string) int {
	count := 0

	for range str {
		count++
	}

	return count
}

func main() {
	// l := Strlen("Hello World!")
	c := Strlen("Hello World!")
	fmt.Println(c)
}
