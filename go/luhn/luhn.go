package luhn

import (
	"strings"
	"unicode"
)

const testVersion = 2

func Valid(input string) bool {
	input = strings.Replace(input, " ", "", -1)
	if len(input) < 2 {
		return false
	}
	rem := len(input) % 2
	sum := 0
	for i, l := range input {
		if !unicode.IsNumber(l) {
			return false
		}
		sum += WeightOfLetter(i, l, rem)
	}

	return sum%10 == 0
}

func WeightOfLetter(i int, l rune, rem int) (val int) {
	if i%2 == rem {
		val = int(l-'0') * 2
		if val > 9 {
			val -= 9
		}
	} else {
		val = int(l - '0')
	}
	return
}
