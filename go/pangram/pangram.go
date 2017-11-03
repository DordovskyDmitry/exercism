package pangram

import "strings"

const testVersion = 2

func IsPangram(s string) bool {
	s = strings.ToLower(s)
	for l := 'a'; l <= 'z'; l++ {
		if !strings.ContainsRune(s, l) {
			return false
		}
	}
	return true
}
