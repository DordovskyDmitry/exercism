package atbash

import (
	"regexp"
	"strings"
	"unicode"
)

func Atbash(in string) string {
	reg, err := regexp.Compile("[^a-z0-9]+")
	if err != nil {
		return ""
	}
	sanitizedIn := reg.ReplaceAllString(strings.ToLower(in), "")

	s := make([]rune, 0)
	pos := 'a' + 'z'
	last := len(sanitizedIn) - 1
	for i, r := range sanitizedIn {
		if unicode.IsDigit(r) {
			s = append(s, r)
		} else {
			s = append(s, pos-r)
		}

		if i%5 == 4 && i != last {
			s = append(s, ' ')
		}
	}
	return string(s)
}
