package main

import "fmt"

func ConcatSlice(slice1, slice2 []int) []int {
	if len(slice1) == 0 && len(slice2) == 0 {
		return nil
	}
	concat := make([]int, len(slice1)+len(slice2))

	copy(concat, slice1)

	copy(concat[len(slice1):], slice2)

	return concat
}

func main() {
	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatSlice([]int{}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{}))
}
