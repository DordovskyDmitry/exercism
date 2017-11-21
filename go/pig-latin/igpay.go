package igpay

import "strings"

var vowels = map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}

func PigLatin(in string) string {
	words := strings.Split(in, " ")
	translates := make([]string, 0, len(words))
	for _, word := range words {
		consonant, other := partition(word)
		translates = append(translates, other+consonant+"ay")
	}

	return strings.Join(translates, " ")
}

func partition(in string) (string, string) {
	i := 0
	if in[0] == 'x' && len(in) > 1 && !vowels[in[1]] {
		return "", in
	}
	for i < len(in) {
		if in[i] == 'y' || vowels[in[i]] {
			break
		}
		i++
	}
	if in[i] == 'y' && i != len(in)-1 && vowels[in[i+1]] {
		return in[:i+1], in[i+1:]
	}
	if in[i] == 'u' && i != 0 && in[i-1] == 'q' {
		return in[:i+1], in[i+1:]
	}
	return in[:i], in[i:]
}
