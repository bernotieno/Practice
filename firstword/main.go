package main

import "fmt"

func FirstWord(str string) string {
	sentence := []rune(str)
	var firstWord string

	// iterate over input string to find the first word
	for i := 0; i < len(sentence); i++ {
		// if space is encountered and word isn't empty, found the first word
		if sentence[i] == ' ' {
			if len(firstWord) > 0 {
				break
			}
		} else {
			firstWord += string(sentence[i])
		}
	}
	return firstWord
}

func main() {
	fmt.Println(FirstWord("hello there"))
	fmt.Println(FirstWord(""))
	fmt.Println(FirstWord("hello   .........  bye"))
}
