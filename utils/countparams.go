package functions

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	isNeg := false
	if n < 0 {
		isNeg = true
		n = -n
	}

	var result []byte
	for n > 0 {
		digit := n % 10
		result = append([]byte{byte('0' + digit)}, result...)
		n /= 10
	}

	if isNeg {
		result = append([]byte{'-'}, result...)
	}
	return string(result)
}

func Countparams() {
	var count int

	arguments := os.Args[1:]

	for range arguments {
		count++
	}

	stri := Itoa(count)

	for _, val := range stri {
		z01.PrintRune(val)
	}
	z01.PrintRune('\n')

	fmt.Println(count)
}
