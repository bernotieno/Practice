package main

import "fmt"

func BeZero(slice []int) []int {
	for i := range slice {
		slice[i] = 0
	}
	return slice
}

func main() {
	fmt.Println(BeZero([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(BeZero([]int{}))
}
