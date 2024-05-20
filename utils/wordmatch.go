package functions

func containWord(first, second string) bool {
	j := 0
	for i := 0; i < len(second) && j < len(first); i++ {
		if first[j] == second[i] {
			j++
		}
	}
	return j == len(first)
}

func WordMatch(first, second string) string {
	if containWord(first, second) {
		return first
	}
	return ""
}

// func Print(str string) {
// 	for _, char := range str {
// 		z01.PrintRune(char)
// 	}
// 	z01.PrintRune('\n')
// }
