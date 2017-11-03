package grains

import "errors"

const testVersion = 1

func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("invalid input")
	}
	m := uint(n - 1)
	return 1 << m, nil
}

func Total() uint64 {
	return (1 << 64) - 1
}
