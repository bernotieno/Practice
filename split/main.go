package main

import "fmt"

func Split(s, sep string) []string {
	var result []string
	start := 0
	for i := 0; i+len(sep) <= len(s); i++ {
		if s[i:i+len(sep)] == sep {
			if start < i {
				result = append(result, s[start:i])
				start = i + len(sep)
			}
		}
	}
	if start < len(s) {
		result = append(result, s[start:])
	}
	return result
}

func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", Split(s, "HA"))
}
