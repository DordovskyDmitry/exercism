package wordcount

import (
	"strings"
	"unicode"
)

const testVersion = 3

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	res := Frequency{}
	i := 0
	for i < len(phrase) {
		if IsAlphaNum(phrase[i]) {
			j := 1
			for i+j < len(phrase) {
				if !(IsAlphaNum(phrase[i+j]) || ContractionQuote(i+j, phrase)) {
					break
				}
				j++
			}
			word := phrase[i : i+j]
			res[strings.ToLower(word)] += 1
			i += j
		}
		i++
	}
	return res
}

func IsAlphaNum(r byte) bool {
	return unicode.IsDigit(rune(r)) || unicode.IsLetter(rune(r))
}

func ContractionQuote(pos int, phrase string) bool {
	return phrase[pos:pos+1] == "'" && pos+1 < len(phrase) && IsAlphaNum(phrase[pos+1])
}
