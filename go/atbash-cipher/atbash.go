package atbash

import "strings"

var mapLetters map[rune]rune = map[rune]rune{}

func init() {
	for r := 'a'; r <= 'z'; r++ {
		mapLetters[r] = 'z' + 'a' - r
	}

	for r := 'A'; r <= 'Z'; r++ {
		mapLetters[r] = 'z' + 'A' - r
	}

	for r := '0'; r <= '9'; r++ {
		mapLetters[r] = r
	}
}

func Atbash(in string) string {
	in = strings.Map(func(r rune) rune {
		if v, ok := mapLetters[r]; ok {
			return v
		}
		return -1
	}, in)
	strs := []string{}
	i := 0
	for i < len(in)-4 {
		strs = append(strs, in[i:i+5])
		i += 5
	}
	if i < len(in) {
		strs = append(strs, in[i:len(in)])
	}
	return strings.Join(strs, " ")
}
