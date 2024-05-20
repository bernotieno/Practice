package functions

func Nrune(s string, n int) rune {
	if n <= 0 || n > len(s) {
		return 0
	}
	return []rune(s)[n-1]
}