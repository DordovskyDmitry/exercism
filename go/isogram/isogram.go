package isogram

import (
	"strings"
	"unicode"
)

const testVersion = 1

func IsIsogram(input string) bool {
	cache := make(map[rune]bool, 26)
	input = strings.ToLower(input)
	for _, l := range input {
		if unicode.IsLetter(l) {
			if cache[l] {
				return false
			}
			cache[l] = true
		}
	}
	return true
}
