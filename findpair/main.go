package main

import (
	"fmt"
	"os"
	"strconv"
)

func split(str, sep string) []int {
	str = str[1:]
	str = str[:len(str)-1]
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
	var res []int
	for _, c := range result {
		n, err := strconv.Atoi(c)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		res = append(res, n)
	}
	return res
}

func pair(slice []int, target int) [][]int {
	var paired [][]int
	for i := 0; i < len(slice)-1; i++ {
		for j := i + 1; j < len(slice); j++ {
			if (slice[i] + slice[j]) == target {
				paired = append(paired, []int{i, j})
			}
		}
	}
	return paired
}

func main() {
	if len(os.Args) != 3 {
		return
	}
	args := os.Args[1]
	target, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}

	pppz := pair(split(args, ", "), target)
	fmt.Println(pppz)
}
