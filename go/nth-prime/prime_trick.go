package prime

var primes = []int{2, 3}

func NthTrick(n int) (int, bool) {
	if n < 1 {
		return 0, false
	}
	i := primes[len(primes)-1] + 2
	for n > len(primes) {
		if isPrimeTrick(i) {
			primes = append(primes, i)
		}
		i += 2

	}
	return primes[n-1], true
}

func isPrimeTrick(n int) bool {
	for _, i := range primes {
		if n%i == 0 {
			return false
		}
	}
	return true
}
