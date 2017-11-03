package palindrome

import (
	"errors"
	"math"
)

const testVersion = 1

type Product struct {
	prod           int
	Factorizations [][2]int
}

func Products(min, max int) (p1 Product, p2 Product, e error) {
	if max < min {
		e = errors.New("fmin > fmax")
		return
	}
	for i := min * min; i <= max*max; i++ {
		if isPalindrome(i) {
			divisors := findDivisors(i, min, max)
			if len(divisors) > 0 {
				p1 = Product{prod: i, Factorizations: divisors}
				break
			}
		}
	}

	for i := max * max; i >= min*min; i-- {
		if isPalindrome(i) {
			divisors := findDivisors(i, min, max)
			if len(divisors) > 0 {
				p2 = Product{prod: i, Factorizations: divisors}
				break
			}
		}
	}

	if p1.prod == 0 && p2.prod == 0 {
		e = errors.New("no palindromes")
		return
	}

	return p1, p2, nil
}

func isPalindrome(i int) bool {
	rev, old := 0, i
	for i > 0 {
		rev = rev*10 + i%10
		i /= 10
	}
	return rev == old
}

func findDivisors(n, min, max int) (res [][2]int) {
	for i := min; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i != 0 {
			continue
		}
		q := n / i
		if q <= max {
			res = append(res, [2]int{i, q})
		}
	}
	return res
}
