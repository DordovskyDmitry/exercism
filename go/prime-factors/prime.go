package prime

func Factors(n int64) (res []int64) {
	return factors(n, 2, make([]int64, 0))
}

func factors(n int64, start int64, divisors []int64) []int64 {
	if n == 1 {
		return divisors
	}
	i := start
	for i*i <= n {
		if n%i == 0 {
			divisors = append(divisors, i)
			return factors(n/i, i, divisors)
		}
		i++
	}
	return append(divisors, n)
}
