package main

import "fmt"

func RemoveDuplicate(slice []int) []int {
	numbMap := make(map[int]bool)
	numbSlice := []int{}

	for _, v := range slice {
		if !numbMap[v] {
			numbMap[v] = true
			numbSlice = append(numbSlice, v)
		}
	}
	return numbSlice
}

func main() {
	fmt.Println(RemoveDuplicate([]int{1, 2, 3, 3, 4, 5, 6, 7, 8, 9, 10}))
	fmt.Println(RemoveDuplicate([]int{1, 1, 2, 2, 3}))
	fmt.Println(RemoveDuplicate([]int{}))
}