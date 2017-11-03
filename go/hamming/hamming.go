package hamming

import "errors"

const testVersion = 6

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Lengths are different")
	}

	s := 0
	for i := 0; i < len(b); i++ {
		if b[i] != a[i] {
			s++
		}
	}
	return s, nil
}
