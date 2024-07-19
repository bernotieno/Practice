package main

import "fmt"

func RevConcatAlternate(slice1, slice2 []int) []int {
	// Reverse the slices
	reversedSlice1 := reverseSlice(slice1)
	reversedSlice2 := reverseSlice(slice2)

	var result []int
	i, j := 0, 0
	len1, len2 := len(reversedSlice1), len(reversedSlice2)

	// Alternate adding elements from both slices, starting with the longer slice
	for i < len1 && j < len2 {
		if len1 > len2 {
			result = append(result, reversedSlice1[i])
			i++
		} else if len2 > len1 {
			result = append(result, reversedSlice2[j])
			j++
		} else {
			result = append(result, reversedSlice1[i])
			i++
		}
		if i < len1 {
			result = append(result, reversedSlice1[i])
			i++
		}
		if j < len2 {
			result = append(result, reversedSlice2[j])
			j++
		}
	}

	// Append remaining elements if any
	for i < len1 {
		result = append(result, reversedSlice1[i])
		i++
	}
	for j < len2 {
		result = append(result, reversedSlice2[j])
		j++
	}

	return result
}

func reverseSlice(slice []int) []int {
	reversed := make([]int, len(slice))
	for i, v := range slice {
		reversed[len(slice)-i-1] = v
	}
	return reversed
}

func main() {
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3, 9, 8}, []int{4, 5}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{}))
}
