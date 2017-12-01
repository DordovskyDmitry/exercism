package prime

func Nth(n int) (int, bool) {
	if n < 1 {
		return 0, false
	}
	if n == 1 {
		return 2, true
	}
	if n == 2 {
		return 3, true
	}
	m := 3
	for n > 2 {
		m += 2
		if millerRabinCheck(m) {
			n--
		}
	}
	return m, true
}

func millerRabinCheck(n int) bool {
	s := 0
	t := n - 1
	for t%2 == 0 {
		t = t / 2
		s++
	}
	for _, a := range []int{2, 3} {
		x := pow(a, t, n)
		if x == 1 || x == n-1 {
			continue
		}
		flag := false
		for j := 0; j < s; j++ {
			x = pow(x, 2, n)
			if x == 1 {
				return false
			}
			if x == n-1 {
				flag = true
				break
			}
		}
		if flag {
			continue
		}
		return false
	}
	return true
}

func pow(a, t, n int) int {
	if t == 1 {
		return a % n
	}
	newA := int(int64(a) * int64(a) % int64(n))
	res := pow(newA, t/2, n)
	if t%2 == 1 {
		res = (res * a) % n
	}
	return res
}
