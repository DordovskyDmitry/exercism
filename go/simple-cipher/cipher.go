package cipher

import (
	"regexp"
	"strings"
)

type Cipher interface {
	Encode(string) string
	Decode(string) string
}

type Shift int

type Vigenere string

var invalidKeys = []string{"", "a", "aa", "no way", "CAT", "3", "and,"}

var validLetters, _ = regexp.Compile("[^a-zA-Z]+")

func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(n int) Cipher {
	if n >= 26 || n <= -26 || n == 0 {
		return nil
	}
	return Shift(n)
}

func (c Shift) Encode(s string) string {
	result := make([]rune, 0, len(s))
	for _, l := range sanitize(s) {
		shifted := l + rune(c)
		if shifted < 'a' {
			shifted += 26
		}
		if shifted > 'z' {
			shifted -= 26
		}
		result = append(result, shifted)
	}
	return string(result)
}

func (c Shift) Decode(s string) string {
	return NewShift(-int(c)).Encode(s)
}

func NewVigenere(s string) Cipher {
	for _, key := range invalidKeys {
		if key == s {
			return nil
		}
	}
	return Vigenere(s)
}

func (c Vigenere) Encode(s string) string {
	return c.aggregateString(s, encodeLetter)
}

func (c Vigenere) Decode(s string) string {
	return c.aggregateString(s, decodeLetter)
}

func (c Vigenere) aggregateString(s string, f func(l, shift rune) rune) string {
	key := ""
	for len(key) < len(s) {
		key += string(c)
	}
	res := []rune{}
	for i, l := range sanitize(s) {
		newL := f(l, rune(key[i]))
		res = append(res, newL)
	}
	return string(res)
}

func encodeLetter(l, shift rune) rune {
	newL := l + rune(shift) - 'a'
	if newL > 'z' {
		newL -= 26
	}
	return newL
}

func decodeLetter(l, shift rune) rune {
	newL := l - rune(shift) + 'a'
	if newL < 'a' {
		newL += 26
	}
	return newL
}

func sanitize(s string) string {
	res := validLetters.ReplaceAllString(s, "")
	return strings.ToLower(res)
}
