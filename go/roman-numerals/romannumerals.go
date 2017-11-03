package romannumerals

import "errors"

const testVersion = 4

func ToRomanNumeral(n int) (string, error) {
	if n > 3000 || n < 1 {
		return "", errors.New("Unsupported number")
	}

	r := []byte{}

	n1000 := n / 1000
	r = appendRest(r, n1000, 'M')

	n100 := n % 1000 / 100
	r = appendNum(r, n100, 'M', 'D', 'C')

	n10 := n % 100 / 10
	r = appendNum(r, n10, 'C', 'L', 'X')

	n1 := n % 10
	r = appendNum(r, n1, 'X', 'V', 'I')

	return string(r), nil
}

func appendNum(to []byte, n int, dec byte, five byte, rest byte) []byte {
	if n == 9 {
		to = append(to, rest, dec)
	} else if n == 4 {
		to = append(to, rest, five)
	} else {
		if n >= 5 {
			to = append(to, five)
			n -= 5
		}
		to = appendRest(to, n, rest)
	}
	return to
}

func appendRest(to []byte, times int, what byte) []byte {
	for i := 0; i < times; i++ {
		to = append(to, what)
	}
	return to
}
