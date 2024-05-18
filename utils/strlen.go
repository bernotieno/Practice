package functions

import "fmt"

func Strlen(s string) {
	count := 0
	for range s {
		count++
	}
	fmt.Println(count)
}
