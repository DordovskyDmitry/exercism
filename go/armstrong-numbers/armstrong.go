package armstrong

func IsNumber(n int) bool {
	digits := make([]int, 0)
	m := n
	for m > 0 {
		digits = append(digits, m%10)
		m = m / 10
	}
	s := 0
	for _, d := range digits {
		s += pow(d, len(digits))
	}
	return s == n
}

func pow(a, t int) int {
	if t == 1 {
		return a
	}
	res := pow(a*a, t/2)
	if t%2 == 1 {
		res *= a
	}
	return res
}
