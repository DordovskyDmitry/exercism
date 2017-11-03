package anagram

import (
	"strings"
)

const testVersion = 2

func Detect(subject string, candidates []string) (res []string) {
	subject = strings.ToLower(subject)
	subjectMap := lettersCount(subject)

	for _, candidate := range candidates {
		candidateLower := strings.ToLower(candidate)
		if len(subject) != len(candidate) || subject == candidateLower {
			continue
		}
		isAnagram := true
		for k, v := range subjectMap {
			if strings.Count(candidateLower, k) != v {
				isAnagram = false
				break
			}
		}
		if isAnagram {
			res = append(res, candidate)
		}
	}
	return
}

func lettersCount(word string) map[string]int {
	res := map[string]int{}
	for i := range word {
		res[word[i:i+1]] += 1
	}
	return res
}
