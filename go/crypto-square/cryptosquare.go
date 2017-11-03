package cryptosquare

import (
	"strings"
	"unicode"
)

const testVersion = 2

func Encode(input string) string {
	s := Sanitize(strings.ToLower(input))
	c, r := Rect(len(s))
	str := make([]string, c)
	for i := 0; i < c; i++ {
		str[i] = ""
		for j := 0; j < r; j++ {
			if i+j*c >= len(s) {
				break
			}
			str[i] += s[i+j*c : i+j*c+1]
		}
	}

	return strings.Join(str, " ")
}

func Sanitize(input string) string {
	sanitized := ""
	for i, l := range input {
		if unicode.IsLetter(l) || unicode.IsNumber(l) {
			sanitized += input[i : i+1]
		}
	}
	return sanitized
}

func Rect(n int) (int, int) {
	for i := 1; i <= n; i++ {
		if i*i >= n {
			return i, i
		}
		if i*(i+1) >= n {
			return i + 1, i
		}
	}
	return 0, 0
}
