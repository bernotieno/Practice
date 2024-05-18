package functions

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func PreviousPrime(nb int) int {
	nb-- // Start checking from the number just below nb
	for nb > 1 {
		if isPrime(nb) {
			return nb
		}
		nb--
	}
	return -1 // Return -1 if no prime number is found
}
