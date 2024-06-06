package main

import "fmt"

func isPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	for i := 2; i*i <= nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func FindPrevPrime(n int) int {
	if n <= 1 {
		return 0
	}

	for {
		if isPrime(n) {
			return n
		}
		n--
	}
}

func main() {
	fmt.Println(FindPrevPrime(5))
	fmt.Println(FindPrevPrime(4))
}
