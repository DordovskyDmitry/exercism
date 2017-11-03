package sieve

const testVersion = 1

func Sieve(n int) (res []int) {
	if n < 2 {
		return res
	}
	notPrimes := make([]bool, n+1, n+1)
	for i := 2; i <= n; i++ {
		if notPrimes[i] {
			continue
		}
		for j := 2 * i; j <= n; j += i {
			notPrimes[j] = true
		}
		if !notPrimes[i] {
			res = append(res, i)
		}
	}
	return
}
