package etl

import "strings"

const testVersion = 1

func Transform(input map[int][]string) (m map[string]int) {
	m = make(map[string]int)
	for key, values := range input {
		for _, v := range values {
			m[strings.ToLower(v)] = key
		}
	}
	return
}
