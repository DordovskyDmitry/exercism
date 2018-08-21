package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	return CollatzConjectureAgg(n, 0, nil)
}

func CollatzConjectureAgg(n int, steps int, err error) (int, error) {
	if err != nil {
		return 0, err
	}
	if n <= 0 {
		return 0, errors.New("non-positive number")
	}
	if n == 1 {
		return steps, err
	}
	var newN int
	if n%2 == 0 {
		newN = n / 2
	} else {
		newN = 3*n + 1
	}
	return CollatzConjectureAgg(newN, steps+1, err)
}
