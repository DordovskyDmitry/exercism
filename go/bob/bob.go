package bob

import "strings"
import "unicode"

const testVersion = 3

func Hey(s string) string {
	s = strings.TrimSpace(s)
	hasLetter := HasLetter(s)

	if hasLetter && AllUpper(s) {
		return "Whoa, chill out!"
	}
	if len(s) > 0 && s[len(s)-1] == '?' {
		return "Sure."
	}
	if !hasLetter && !HasDigit(s) {
		return "Fine. Be that way!"
	}
	return "Whatever."
}

func HasLetter(s string) bool {
	hasLetter := false
	for _, l := range s {
		if unicode.IsLetter(l) {
			hasLetter = true
			break
		}
	}
	return hasLetter
}

func AllUpper(s string) bool {
	hasLower := false
	for _, l := range s {
		if l >= 'a' && l <= 'z' {
			hasLower = true
			break
		}
	}
	return !hasLower
}

func HasDigit(s string) bool {
	hasDigit := false
	for _, l := range s {
		if l >= '0' && l <= '9' {
			hasDigit = true
			break
		}
	}
	return hasDigit
}
